package swap

import (
	"time" // Para time.Now() em SetApprovedAt

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ss *swapDomainService) UpdateSwapServices(
	id string, // ID da Swap a ser atualizada
	swapUpdateDomain model.SwapDomainInterface, // Contém os dados para atualização (pode ser um status, approvedBy, etc.)
) *rest_err.RestErr {
	logger.Info("Init UpdateSwapServices", // Nome da função no log
		zap.String("journey", "updateSwap"),
		zap.String("swapID", id))

	if id == "" {
		logger.Error("Swap ID for UpdateSwapServices cannot be empty", nil,
			zap.String("journey", "updateSwap"))
		return rest_err.NewBadRequestError("Swap ID cannot be empty")
	}

	// 1. Buscar a Swap existente para garantir que ela existe antes de atualizar.
	existingSwap, findErr := ss.repository.FindSwapByID(id)
	if findErr != nil {
		logger.Error("Swap to update not found by service", findErr,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return findErr // Retorna o erro do repositório (ex: NotFoundError)
	}

	// 2. Aplicar as atualizações do swapUpdateDomain ao existingSwap.
	// O controller é responsável por construir o swapUpdateDomain corretamente
	// com base no que pode ser atualizado (ex: apenas o status e quem aprovou).
	// Exemplo: Atualizando status e dados de aprovação.
	// A sua struct swapDomain já tem SetStatus, SetApprovedAt, SetApprovedBy.
	// O seu NewSwapUpdateDomain não tem campos para status, approvedAt, approvedBy.
	// O controller chama NewSwapUpdateDomain e depois SetApprovedBy.
	// O status vem da request e é setado no domain pelo controller antes de chamar este serviço.

	// O swapUpdateDomain que chega aqui (vindo do controller) já tem:
	// - Os campos de Current/New Shift/DayOff (do NewSwapUpdateDomain).
	// - ApprovedBy (setado pelo controller).
	// - O Status deve ser setado no controller com base na request /status ANTES de chamar este serviço.

	// O que precisamos fazer é transferir os campos mutáveis de swapUpdateDomain para existingSwap
	// ou, se swapUpdateDomain já é o estado final desejado (exceto ID e CreatedAt),
	// podemos passá-lo diretamente para o repositório.
	// A sua lógica atual no controller de UpdateSwapStatus parece criar um novo domain
	// e então chamar o serviço de update.

	// Vamos assumir que swapUpdateDomain contém os campos que DEVEM ser atualizados.
	// E que `existingSwap` é a nossa base.

	// Exemplo de lógica de atualização no serviço:
	// O controller de UpdateSwapStatus recebe um `status` (pending, approved, rejected).
	// Ele deve construir um `swapUpdateDomain` (usando `NewSwapUpdateDomain` ou similar)
	// e então explicitamente setar o status, e se for 'approved', o `approvedAt` e `approvedBy`.

	// Se o status está mudando para "approved":
	if swapUpdateDomain.GetStatus() == model.StatusApproved && existingSwap.GetStatus() != model.StatusApproved {
		// Garante que approvedAt e approvedBy sejam definidos se o status for 'approved'.
		// O controller já seta o ApprovedBy.
		// Aqui podemos setar o ApprovedAt.
		if swapUpdateDomain.GetApprovedAt() == nil || swapUpdateDomain.GetApprovedAt().IsZero() { // Se o controller não setou
			now := time.Now()
			swapUpdateDomain.SetApprovedAt(now) // O SetApprovedAt na sua interface usa ponteiro, precisa ajustar
		}
	} else if swapUpdateDomain.GetStatus() == model.StatusRejected && existingSwap.GetStatus() != model.StatusRejected {
		// Se rejeitado, talvez limpar ApprovedAt e ApprovedBy?
		// A sua interface SwapDomainInterface não tem métodos para "limpar" (setar para nil) ApprovedAt/By.
		// Precisaria adicionar `ClearApprovedAt()` e `ClearApprovedBy()` ou permitir `SetApprovedAt(nil)`.
		// Por agora, o `updateFields` no repositório setará para o que estiver no domain.
	} else if swapUpdateDomain.GetStatus() == model.StatusPending && existingSwap.GetStatus() != model.StatusPending {
		// Se voltando para pendente, limpar campos de aprovação.
	}

	// O repositório UpdateSwap já constrói o bson.M com os campos do domain passado.
	// Passamos o swapUpdateDomain que contém as informações atualizadas + as originais não modificadas.
	// No entanto, o repositório UpdateSwap que projetei espera todos os campos, o que pode
	// sobrescrever campos não intencionalmente se swapUpdateDomain não for o estado completo.
	// A melhor abordagem é o repositório receber um domain que é o estado *final* desejado,
	// e ele faz um $set com todos os campos desse domain (exceto _id e talvez createdAt).

	// Vamos refinar: o serviço deve preparar o *estado final completo* do `existingSwap`
	// e então passar esse `existingSwap` (modificado) para o repositório.

	// Atualiza os campos de existingSwap com base no que veio em swapUpdateDomain
	// (que foi preparado pelo controller)
	existingSwap.SetStatus(swapUpdateDomain.GetStatus())
	if swapUpdateDomain.GetApprovedBy() != nil && *swapUpdateDomain.GetApprovedBy() != "" {
		existingSwap.SetApprovedBy(*swapUpdateDomain.GetApprovedBy())
	}
	if swapUpdateDomain.GetApprovedAt() != nil && !swapUpdateDomain.GetApprovedAt().IsZero() {
		existingSwap.SetApprovedAt(*swapUpdateDomain.GetApprovedAt())
	} else if swapUpdateDomain.GetStatus() == model.StatusApproved && (existingSwap.GetApprovedAt() == nil || existingSwap.GetApprovedAt().IsZero()) {
		// Se o status é approved e approvedAt não foi explicitamente setado pelo controller (via swapUpdateDomain),
		// então setamos agora.
		now := time.Now()
		existingSwap.SetApprovedAt(now)
	}

	// Os outros campos como reason, shifts, daysoff, vêm do swapUpdateDomain,
	// que no seu controller é construído com NewSwapUpdateDomain.
	// Se a intenção é que esses também possam ser atualizados pela rota /status,
	// o request precisaria incluí-los. Assumindo que a rota /status só muda o status e aprovação.
	// No entanto, seu NewSwapUpdateDomain inclui shifts e daysoff.
	// Se o UpdateSwapServices for genérico para qualquer update de Swap:
	existingSwap.SetRequestedID(swapUpdateDomain.GetRequestedID()) // Se puder mudar
	existingSwap.SetCurrentShift(swapUpdateDomain.GetCurrentShift())
	existingSwap.SetNewShift(swapUpdateDomain.GetNewShift())
	existingSwap.SetCurrentDayOff(swapUpdateDomain.GetCurrentDayOff())
	existingSwap.SetNewDayOff(swapUpdateDomain.GetNewDayOff())
	existingSwap.SetReason(swapUpdateDomain.GetReason())
	// RequesterID e CreatedAt não devem mudar em um update.

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
