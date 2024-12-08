package service

import (
	"time"

	"automatic-doodle/types"
)

func (svc *Service) Initialize() error {
	svc.customerCreds = types.JWTCredentials{}
	var err error
	if svc.customerCreds.Aud, err = svc.cfg.GetConfig(
		"CUSTOMER_TOKEN_JWT_AUDIANCE",
	); err != nil {
		svc.log.Fatal("Missing JWT Token env variable: %s", err)
		return err
	}
	if svc.customerCreds.Aud, err = svc.cfg.GetConfig(
		"CUSTOMER_TOKEN_JWT_SECRET_KEY",
	); err != nil {
		svc.log.Fatal("Missing JWT Token env variable: %s", err)
		return err
	}
	var customerExpMin int
	if customerExpMin, err = svc.cfg.GetConfigInt(
		"CUSTOMER_TOKEN_JWT_EXPIRE_MIN",
	); err != nil {
		svc.log.Fatal("Missing JWT Token env variable: %s", err)
		return err
	}
	svc.customerCreds.ExpireDuration = time.Duration(customerExpMin) * time.Minute

	svc.adminCreds = types.JWTCredentials{}
	if svc.adminCreds.Aud, err = svc.cfg.GetConfig(
		"ADMIN_TOKEN_JWT_AUDIANCE",
	); err != nil {
		svc.log.Fatal("Missing JWT Token env variable: %s", err)
		return err
	}
	if svc.adminCreds.Aud, err = svc.cfg.GetConfig(
		"ADMIN_TOKEN_JWT_SECRET_KEY",
	); err != nil {
		svc.log.Fatal("Missing JWT Token env variable: %s", err)
		return err
	}
	var adminExpMin int
	if adminExpMin, err = svc.cfg.GetConfigInt(
		"ADMIN_TOKEN_JWT_EXPIRE_MIN",
	); err != nil {
		svc.log.Fatal("Missing JWT Token env variable: %s", err)
		return err
	}
	svc.adminCreds.ExpireDuration = time.Duration(adminExpMin) * time.Minute

	return nil
}
