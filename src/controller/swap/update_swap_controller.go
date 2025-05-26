package swap

import (
	"net/http"
	"time"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/validation"

	// Import para o DTO de request de swap, usando alias
	swap_request_dto "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/swap/request"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive" // Para validar o formato do ID
	"go.uber.org/zap"
)

func (sc *swapControllerInterface) UpdateSwapStatus(c *gin.Context) {
	logger.Info("Init UpdateSwapStatus controller",
		zap.String("journey", "updateSwapStatus"))

	swapID := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(swapID); err != nil {
		logger.Error("Invalid swap ID format in UpdateSwapStatus controller", err,
			zap.String("journey", "updateSwapStatus"),
			zap.String("swapID", swapID))
		restErrVal := rest_err.NewBadRequestError("Invalid Swap ID format, must be a hex value.")
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	var statusRequest swap_request_dto.SwapRequest // Usando o DTO flexível
	if err := c.ShouldBindJSON(&statusRequest); err != nil {
		logger.Error("Error validating status update request data", err,
			zap.String("journey", "updateSwapStatus"))
		restErrVal := validation.ValidateUserError(err)
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	// Validação para garantir que o campo status está presente para atualização de status
	if statusRequest.Status == "" {
		logger.Error("Missing status field for swap status update", nil,
			zap.String("journey", "updateSwapStatus"))
		restErrVal := rest_err.NewBadRequestError("Missing status field for swap status update")
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	newStatus := model.SwapStatus(statusRequest.Status)
	if newStatus != model.StatusApproved && newStatus != model.StatusRejected && newStatus != model.StatusPending {
		logger.Error("Invalid status value for swap status update", nil,
			zap.String("journey", "updateSwapStatus"),
			zap.String("receivedStatus", statusRequest.Status))
		restErrVal := rest_err.NewBadRequestError("Invalid status value. Must be 'approved', 'rejected', or 'pending'.")
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	// TODO: (Pós-JWT) Lógica de Permissão:
	// - Quem pode aprovar/rejeitar? (Master? Superior do requestedID? Superior do requesterID?)
	// - O 'approverID' deve vir do token JWT.
	// approverID := c.GetString("userID") // Exemplo
	approverID := "temp-approver-id" // Placeholder - REMOVER/AJUSTAR COM JWT
	if approverID == "" {
		logger.Error("Approver ID not found (simulate JWT)", nil, zap.String("journey", "updateSwapStatus"))
		restErr_ := rest_err.NewUnauthorizedError("Unauthorized: Approver ID not found.") // Corrigido nome da var
		c.JSON(restErr_.Code, restErr_)
		return
	}

	// Para o UpdateSwapServices, precisamos de um SwapDomainInterface.
	// O NewSwapUpdateDomain no seu código original era usado para construir um domain
	// a partir de uma request que parecia ser de criação, não de atualização de status.
	// Para atualizar o status, precisamos apenas do novo status e quem aprovou/rejeitou.
	// O serviço de update agora busca o swap existente e aplica as mudanças.

	// Vamos criar um domain apenas com as informações relevantes para a atualização de status.
	// O serviço de update irá carregar o swap existente e aplicar estas mudanças.
	// O NewSwapUpdateDomain que você tinha pode ser adaptado ou podemos criar um mais específico.
	// Por agora, vamos passar um domain que o serviço usará para extrair Status, ApprovedBy, ApprovedAt.

	// O seu service.UpdateSwapServices agora espera um model.SwapDomainInterface
	// com os campos que devem ser atualizados.
	// Vamos construir um domain apenas com o necessário para o update de status.
	// O NewSwapUpdateDomain não é ideal aqui pois ele seta shifts e daysoff que não vêm na request de /status.
	// Vamos construir um domain mais simples ou ajustar o serviço para pegar apenas status e approvedBy/At.

	// A forma como o UpdateSwapServices foi ajustado é: ele carrega o `existingSwap`
	// e depois aplica os campos de um `swapUpdateDomain` nele.
	// Então, precisamos criar um `swapUpdateDomain` que só tenha os campos que queremos mudar:
	// Status, ApprovedBy, e ApprovedAt (se aplicável).

	// Criamos um domain "parcial" ou "delta" para o update.
	// O NewSwapDomain pode ser usado com campos vazios para o que não muda.
	// Mas isso não é o ideal. Melhor seria ter um domain específico para update de status,
	// ou o serviço ser mais inteligente.

	// Dado o UpdateSwapServices atual, ele espera que o domain passado tenha os campos
	// de Status, ApprovedBy, ApprovedAt (se status for approved).

	// O controller precisa preparar esse `swapUpdateDomain`.
	// O `swapRequest.RequestedID` e outros campos de shift/dayoff vindos do `ShouldBindJSON`
	// com `SwapRequest` seriam vazios aqui se o cliente só enviou `status`.
	// O `NewSwapDomain` espera todos esses campos.

	// Vamos simplificar: o `UpdateSwapServices` agora pega o ID e um `model.SwapDomainInterface`
	// que representa o estado *a ser atualizado*.
	// O controller deve construir este domain.

	// Recuperar o Swap existente para obter os dados que não mudam.
	// Esta é uma responsabilidade que o serviço já está fazendo (FindSwapByID dentro de UpdateSwapServices).
	// Então, o controller só precisa preparar o domain com os dados que *mudam*.

	updatePayload := model.NewSwapDomain("", "", "", "", "", "", "") // Campos não relevantes vazios

	updatePayload.SetStatus(newStatus)
	if newStatus == model.StatusApproved {
		updatePayload.SetApprovedBy(approverID)
		now := time.Now()
		updatePayload.SetApprovedAt(now) // O serviço também pode fazer isso
	} else if newStatus == model.StatusRejected {
		// Se rejeitado, e se ApprovedBy deve ser registrado como quem rejeitou.
		// updatePayload.SetApprovedBy(approverID) // Opcional, depende da sua lógica de negócio
		// ApprovedAt permaneceria nil.
	}

	serviceErr := sc.service.UpdateSwapServices(swapID, updatePayload)
	if serviceErr != nil {
		logger.Error("Failed to call swap status update service", serviceErr,
			zap.String("journey", "updateSwapStatus"),
			zap.String("swapID", swapID))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("UpdateSwapStatus controller executed successfully",
		zap.String("swapID", swapID),
		zap.String("newStatus", string(newStatus)),
		zap.String("journey", "updateSwapStatus"))

	c.Status(http.StatusOK)
}
