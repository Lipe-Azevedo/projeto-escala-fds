package comment

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

// ConvertCommentEntityToDomain converte um CommentEntity para CommentDomainInterface.
func ConvertCommentEntityToDomain(
	commentEntity entity.CommentEntity,
) domain.CommentDomainInterface {
	// Cria o domínio usando o construtor que já normaliza a data se necessário,
	// embora aqui a data já venha do BD.
	// No entanto, para manter consistência, podemos usar NewCommentDomain e depois setar os outros campos.
	// Ou, ter um construtor de hidratação mais direto no domínio.
	// Por simplicidade agora, vamos popular diretamente e depois setar ID e timestamps.

	// O construtor NewCommentDomain define createdAt como time.Now(), o que não é desejado para hidratação.
	// Criaremos a instância e setaremos os campos manualmente, respeitando a interface.
	// Esta é uma limitação se o struct `commentDomain` não for exportado ou não tiver um construtor de hidratação.
	// Assumindo que podemos fazer isso através da interface (SetID, SetUpdatedAt).

	domainComment := domain.NewCommentDomain(
		commentEntity.CollaboratorID,
		commentEntity.AuthorID,
		commentEntity.Date, // A data do BD
		commentEntity.Text,
	)

	domainComment.SetID(commentEntity.ID.Hex())

	// Precisamos de uma forma de setar CreatedAt e UpdatedAt no domínio se forem diferentes
	// do que NewCommentDomain define, caso o construtor não seja adequado para hidratação.
	// Se commentDomain fosse um struct local, poderíamos:
	// actualDomain := domainComment.(*domain.commentDomain) // Type assertion
	// actualDomain.createdAt = commentEntity.CreatedAt      // Definir o createdAt do BD
	// actualDomain.updatedAt = commentEntity.UpdatedAt      // Definir o updatedAt do BD

	// Como não podemos fazer isso diretamente sem alterar o domain.go para expor
	// ou ter um construtor de hidratação, a CreatedAt do domainComment será a do momento da conversão.
	// Isto é um ponto de atenção similar ao do SwapDomain.CreatedAt.

	if commentEntity.UpdatedAt != nil {
		domainComment.SetUpdatedAt(*commentEntity.UpdatedAt)
	}
	// Para CreatedAt, o ideal seria ter um domain.NewCommentFromEntity ou similar.
	// Ou, se `commentDomain` fosse exportado:
	// return &domain.commentDomain {
	//  id: commentEntity.ID.Hex(),
	//  collaboratorID: commentEntity.CollaboratorID,
	//  ...
	//  createdAt: commentEntity.CreatedAt, // << Preservaria o CreatedAt do BD
	//  updatedAt: commentEntity.UpdatedAt,
	// }
	// Por ora, o CreatedAt do domínio será o do momento da chamada a NewCommentDomain.

	return domainComment
}
