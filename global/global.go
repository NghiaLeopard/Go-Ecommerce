package global

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/gmail"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/logger"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
)

var (
	Config config.Config
	Logger *logger.LoggerZap
	DB     *db.Queries
	Gmail  gmail.Sender
	Token  token.Maker
)
