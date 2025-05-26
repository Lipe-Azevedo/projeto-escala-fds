package swap

import (
	// "fmt" // Não é mais necessário aqui se não usarmos fmt.Sprintf para erros específicos
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"go.uber.org/zap"
)

func (ss *swapDomainService) UpdateSwapServices(
	id string,
	swapUpdateInfo model.SwapDomainInterface, // Renomeado para clareza, contém info do update (status, approver)
) *rest_err.RestErr {
	logger.Info("Init UpdateSwapServices",
		zap.String("journey", "updateSwap"),
		zap.String("swapID", id))

	if id == "" {
		logger.Error("Swap ID for UpdateSwapServices cannot be empty", nil,
			zap.String("journey", "updateSwap"))
		return rest_err.NewBadRequestError("Swap ID cannot be empty")
	}

	existingSwap, findErr := ss.repository.FindSwapByID(id)
	if findErr != nil {
		logger.Error("Swap to update not found by service", findErr,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return findErr
	}

	// Aplicar apenas as atualizações de status e aprovação
	newStatus := swapUpdateInfo.GetStatus()
	existingSwap.SetStatus(newStatus)

	if newStatus == model.StatusApproved {
		// ApprovedBy é setado pelo controller no swapUpdateInfo
		if swapUpdateInfo.GetApprovedBy() != nil && *swapUpdateInfo.GetApprovedBy() != "" {
			existingSwap.SetApprovedBy(*swapUpdateInfo.GetApprovedBy())
		}
		// ApprovedAt é setado pelo controller no swapUpdateInfo, ou aqui se não estiver presente
		if swapUpdateInfo.GetApprovedAt() != nil && !swapUpdateInfo.GetApprovedAt().IsZero() {
			existingSwap.SetApprovedAt(*swapUpdateInfo.GetApprovedAt())
		} else { // Se o controller não setou ApprovedAt mas o status é approved, setamos agora.
			now := time.Now()
			existingSwap.SetApprovedAt(now)
		}
	} else {
		// Se o status não for "approved" (ex: rejected, pending),
		// podemos querer limpar ApprovedAt e ApprovedBy.
		// A interface SwapDomainInterface não tem métodos para setar ApprovedAt/By para nil.
		// Se precisarmos disso, teríamos que adicionar métodos como ClearApprovedAt() ou modificar os setters
		// para aceitar um ponteiro de tempo ou uma flag.
		// Por agora, o comportamento é que, se não for approved, esses campos podem manter valores antigos
		// se não forem explicitamente alterados no swapUpdateInfo.
		// No entanto, o UpdateSwap do repositório usa $set com o que está no domain.
		// Se o swapUpdateInfo não tiver ApprovedAt/By, e o existingSwap for passado para o repo,
		// os valores antigos de ApprovedAt/By do existingSwap seriam mantidos, A MENOS QUE
		// o swapUpdateInfo (construído com NewSwapDomain e apenas status setado) tivesse
		// ApprovedAt/By como nil/vazio e esses fossem copiados para existingSwap.
		// A lógica atual do controller para `updatePayload` quando não é approved
		// não seta ApprovedAt/By, então o `GetApprovedAt/By` de `swapUpdateInfo` retornaria nil.
		// Se chamarmos `existingSwap.SetApprovedAt(*swapUpdateInfo.GetApprovedAt())` quando é nil, daria panic.
		// Então, a lógica de limpar explicitamente ou não esses campos precisa ser definida.
		// Assumindo que para "rejected" ou "pending", ApprovedAt e ApprovedBy devem ser nulos/vazios:
		// Esta parte requer que a struct swapDomain por trás da interface possa ter esses campos como nil.
		// E os setters na interface teriam que ser capazes de lidar com isso, ou teríamos `Clear` methods.
		// Por simplicidade, vamos assumir que o repositório ao fazer $set com existingSwap, se os campos
		// ApprovedAt e ApprovedBy em existingSwap forem nil (após esta lógica), eles serão persistidos como null.
		// Para fazer isso, precisaríamos de uma forma de setá-los para nil na implementação de swapDomain.
		// A sua interface SetApprovedAt(time.Time) não permite passar nil. SetApprovedBy(string) pode passar string vazia.

		// Vamos manter simples: o controller prepara o swapUpdateInfo. Se lá ApprovedAt/By forem nil,
		// e se os setters permitissem, poderiam ser "limpos".
		// No cenário atual, o mais seguro é não tentar "limpar" explicitamente aqui sem
		// modificar a interface/implementação do domain.
		// Apenas o status é alterado e, se approved, os campos de aprovação são preenchidos.
	}

	// Os campos como RequestedID, shifts, daysoff, reason NÃO SÃO ALTERADOS por este serviço/endpoint.
	// Eles permanecem como estão em existingSwap.

	err := ss.repository.UpdateSwap(id, existingSwap) // Passa o existingSwap modificado
	if err != nil {
		logger.Error("Error calling repository to update swap", err,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return err
	}

	logger.Info("UpdateSwapServices executed successfully",
		zap.String("swapID", id),
		zap.String("journey", "updateSwap"))

	return nil
}
