package loop

import (
	"context"
	"time"

	"github.com/denis96z/simple-version-tracker/worker/pkg/checker"
	"github.com/denis96z/simple-version-tracker/worker/pkg/logs"
	"github.com/denis96z/simple-version-tracker/worker/pkg/storage"
)

type Loop struct {
	conf Config

	strg *storage.Storage
	chkr *checker.Checker
}

func NewLoop(
	conf Config, strg *storage.Storage, chkr *checker.Checker,
) *Loop {
	return &Loop{
		conf: conf,
		strg: strg,
		chkr: chkr,
	}
}

func (lp *Loop) Run(ctx context.Context) error {
	for {
		logs.DebugF("check started")

		arr, err := lp.strg.SelectExternalProjectsToBeCheckedForUpdate(ctx)
		if err != nil {
			logs.Error(err)
			goto next
		}
		logs.DebugF(
			"%d external projects versions will be checked", len(arr),
		)
		for _, item := range arr {
			v, err := lp.chkr.GetLatestVersion(ctx, item.CheckerImage)
			if err != nil {
				logs.Error(err)
				continue
			}
			if v == item.LatestVersion {
				logs.DebugF("%q: no update")
				continue
			}

			// TODO
		}

	next:
		logs.DebugF("CHECK FINISHED")
		time.Sleep(lp.conf.Sleep)
	}
}
