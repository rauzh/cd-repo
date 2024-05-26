package tests

import (
	"context"
	"testing"

	cdtime "github.com/rauzh/cd-core/time"
	"github.com/rauzh/cd-repo/integration_tests/containers"

	"github.com/rauzh/cd-core/requests/base"
	"github.com/rauzh/cd-core/requests/sign_contract"
	signReqPgRepo "github.com/rauzh/cd-repo/pg"
	"github.com/stretchr/testify/assert"
)

func TestSignContractRequest_CreateGet(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	signreqrepo := signReqPgRepo.NewSignContractRequestPgRepo(db)

	ctx := context.Background()

	signReq := sign_contract.SignContractRequest{
		Request: base.Request{
			Type:      sign_contract.SignRequest,
			Status:    base.OnApprovalRequest,
			ApplierID: 2,
			ManagerID: 1,
			Date:      cdtime.GetToday(),
		},
		Nickname:    "zeliboba",
		Description: "",
	}

	err = signreqrepo.Create(ctx, &signReq)

	assert.Equal(t, nil, err)

	signReqCopy, err := signreqrepo.Get(ctx, signReq.RequestID)

	assert.Equal(t, signReqCopy.ManagerID, signReq.ManagerID)
	assert.Equal(t, signReqCopy.RequestID, signReq.RequestID)
	assert.Equal(t, signReqCopy.Type, signReq.Type)
	assert.Equal(t, signReqCopy.ApplierID, signReq.ApplierID)
	assert.Equal(t, signReqCopy.Date, signReq.Date)
	assert.Equal(t, signReqCopy.Nickname, signReq.Nickname)
	assert.Equal(t, signReqCopy.Description, signReq.Description)
}
