package helpers

import (
	"context"
	"piepay/config"
	"piepay/services/es"
	"piepay/services/logger"
	"piepay/structs"
	"piepay/structs/requests"
	"piepay/structs/response"

	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
)

func GetSearchVideo(ctx context.Context, request *requests.SearchVideo, sentryCtx context.Context) (response.GetVideo, error) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] SearchVideo")
	defer span.Finish()

	var res *elastic.SearchResult
	var err error
	if request.Page == 0 {
		request.Size = 10
	}

	index := config.Get().Index

	if len(request.Description) != 0 { //if searching on basis of description,search description in sources
		dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Search from video index")
		res, err = es.Client().Search().Index(index).SearchSource(elastic.NewSearchSource().Query(QueryDetails("description", request.Description)).From(request.Page).Size(request.Size)).Do(ctx)

		dbSpan1.Finish()

	} else { //if searching on basis of title,search title in sources
		dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Get from videos")
		res, err = es.Client().Search().Index(index).SearchSource(elastic.NewSearchSource().Query(QueryDetails("title", request.Title)).SortBy(SortDetails("publishedAt")).From(request.Page).Size(request.Size)).Do(ctx)

		dbSpan1.Finish()

	}

	if err != nil {
		sentry.CaptureException(err)
		logger.Client().Error(err.Error())
		return response.GetVideo{}, err
	}
	var data1 structs.Video
	var dataRes []structs.Video
	if res != nil {
		for _, s := range res.Hits.Hits {
			jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(s.Source, &data1)
			dataRes = append(dataRes, data1)
		}
	}

	getRes := response.GetVideo{
		VideoDetails: dataRes,
		Page:         request.Page + 1,
		Size:         request.Size,
	}
	return getRes, nil
}
