package spam

import (
	"duck/internal/models"
	"duck/internal/pkg/errs"
	"duck/internal/services"
)

type EmailVerifyStrategy struct{}

func (EmailVerifyStrategy) Name() string {
	return "EmailVerifyStrategy"
}

func (EmailVerifyStrategy) CheckTopic(user *models.User, form models.CreateTopicForm) error {
	if services.SysConfigService.IsCreateTopicEmailVerified() && !user.EmailVerified {
		return errs.EmailNotVerified
	}
	return nil
}

func (EmailVerifyStrategy) CheckArticle(user *models.User, form models.CreateArticleForm) error {
	if services.SysConfigService.IsCreateArticleEmailVerified() && !user.EmailVerified {
		return errs.EmailNotVerified
	}
	return nil
}

func (EmailVerifyStrategy) CheckComment(user *models.User, form models.CreateCommentForm) error {
	if services.SysConfigService.IsCreateCommentEmailVerified() && !user.EmailVerified {
		return errs.EmailNotVerified
	}
	return nil
}
