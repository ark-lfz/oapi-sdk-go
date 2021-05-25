// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/api/core/response"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
)

type Service struct {
	conf         *config.Config
	Buildings    *BuildingService
	Countrys     *CountryService
	Districts    *DistrictService
	Freebusys    *FreebusyService
	Instances    *InstanceService
	Rooms        *RoomService
	Summarys     *SummaryService
	MeetingRooms *MeetingRoomService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Buildings = newBuildingService(s)
	s.Countrys = newCountryService(s)
	s.Districts = newDistrictService(s)
	s.Freebusys = newFreebusyService(s)
	s.Instances = newInstanceService(s)
	s.Rooms = newRoomService(s)
	s.Summarys = newSummaryService(s)
	s.MeetingRooms = newMeetingRoomService(s)
	return s
}

type BuildingService struct {
	service *Service
}

func newBuildingService(service *Service) *BuildingService {
	return &BuildingService{
		service: service,
	}
}

type CountryService struct {
	service *Service
}

func newCountryService(service *Service) *CountryService {
	return &CountryService{
		service: service,
	}
}

type DistrictService struct {
	service *Service
}

func newDistrictService(service *Service) *DistrictService {
	return &DistrictService{
		service: service,
	}
}

type FreebusyService struct {
	service *Service
}

func newFreebusyService(service *Service) *FreebusyService {
	return &FreebusyService{
		service: service,
	}
}

type InstanceService struct {
	service *Service
}

func newInstanceService(service *Service) *InstanceService {
	return &InstanceService{
		service: service,
	}
}

type RoomService struct {
	service *Service
}

func newRoomService(service *Service) *RoomService {
	return &RoomService{
		service: service,
	}
}

type SummaryService struct {
	service *Service
}

func newSummaryService(service *Service) *SummaryService {
	return &SummaryService{
		service: service,
	}
}

type MeetingRoomService struct {
	service *Service
}

func newMeetingRoomService(service *Service) *MeetingRoomService {
	return &MeetingRoomService{
		service: service,
	}
}

type SummaryBatchGetReqCall struct {
	ctx      *core.Context
	summarys *SummaryService
	body     *SummaryBatchGetReqBody
	optFns   []request.OptFn
}

func (rc *SummaryBatchGetReqCall) Do() (*SummaryBatchGetResult, error) {
	var result = &SummaryBatchGetResult{}
	req := request.NewRequest("meeting_room/summary/batch_get", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.summarys.service.conf, req)
	return result, err
}

func (summarys *SummaryService) BatchGet(ctx *core.Context, body *SummaryBatchGetReqBody, optFns ...request.OptFn) *SummaryBatchGetReqCall {
	return &SummaryBatchGetReqCall{
		ctx:      ctx,
		summarys: summarys,
		body:     body,
		optFns:   optFns,
	}
}

type FreebusyBatchGetReqCall struct {
	ctx         *core.Context
	freebusys   *FreebusyService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *FreebusyBatchGetReqCall) SetRoomIds(roomIds ...string) {
	rc.queryParams["room_ids"] = roomIds
}
func (rc *FreebusyBatchGetReqCall) SetTimeMin(timeMin string) {
	rc.queryParams["time_min"] = timeMin
}
func (rc *FreebusyBatchGetReqCall) SetTimeMax(timeMax string) {
	rc.queryParams["time_max"] = timeMax
}

func (rc *FreebusyBatchGetReqCall) Do() (*FreebusyBatchGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &FreebusyBatchGetResult{}
	req := request.NewRequest("meeting_room/freebusy/batch_get", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.freebusys.service.conf, req)
	return result, err
}

func (freebusys *FreebusyService) BatchGet(ctx *core.Context, optFns ...request.OptFn) *FreebusyBatchGetReqCall {
	return &FreebusyBatchGetReqCall{
		ctx:         ctx,
		freebusys:   freebusys,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type BuildingBatchGetReqCall struct {
	ctx         *core.Context
	buildings   *BuildingService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *BuildingBatchGetReqCall) SetBuildingIds(buildingIds ...string) {
	rc.queryParams["building_ids"] = buildingIds
}
func (rc *BuildingBatchGetReqCall) SetFields(fields string) {
	rc.queryParams["fields"] = fields
}

func (rc *BuildingBatchGetReqCall) Do() (*BuildingBatchGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &BuildingBatchGetResult{}
	req := request.NewRequest("meeting_room/building/batch_get", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.buildings.service.conf, req)
	return result, err
}

func (buildings *BuildingService) BatchGet(ctx *core.Context, optFns ...request.OptFn) *BuildingBatchGetReqCall {
	return &BuildingBatchGetReqCall{
		ctx:         ctx,
		buildings:   buildings,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type RoomBatchGetReqCall struct {
	ctx         *core.Context
	rooms       *RoomService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *RoomBatchGetReqCall) SetRoomIds(roomIds ...string) {
	rc.queryParams["room_ids"] = roomIds
}
func (rc *RoomBatchGetReqCall) SetFields(fields string) {
	rc.queryParams["fields"] = fields
}

func (rc *RoomBatchGetReqCall) Do() (*RoomBatchGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &RoomBatchGetResult{}
	req := request.NewRequest("meeting_room/room/batch_get", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.rooms.service.conf, req)
	return result, err
}

func (rooms *RoomService) BatchGet(ctx *core.Context, optFns ...request.OptFn) *RoomBatchGetReqCall {
	return &RoomBatchGetReqCall{
		ctx:         ctx,
		rooms:       rooms,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type BuildingBatchGetIdReqCall struct {
	ctx         *core.Context
	buildings   *BuildingService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *BuildingBatchGetIdReqCall) SetCustomBuildingIds(customBuildingIds ...string) {
	rc.queryParams["custom_building_ids"] = customBuildingIds
}

func (rc *BuildingBatchGetIdReqCall) Do() (*BuildingBatchGetIdResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &BuildingBatchGetIdResult{}
	req := request.NewRequest("meeting_room/building/batch_get_id", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.buildings.service.conf, req)
	return result, err
}

func (buildings *BuildingService) BatchGetId(ctx *core.Context, optFns ...request.OptFn) *BuildingBatchGetIdReqCall {
	return &BuildingBatchGetIdReqCall{
		ctx:         ctx,
		buildings:   buildings,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type RoomBatchGetIdReqCall struct {
	ctx         *core.Context
	rooms       *RoomService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *RoomBatchGetIdReqCall) SetCustomRoomIds(customRoomIds ...string) {
	rc.queryParams["custom_room_ids"] = customRoomIds
}

func (rc *RoomBatchGetIdReqCall) Do() (*RoomBatchGetIdResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &RoomBatchGetIdResult{}
	req := request.NewRequest("meeting_room/room/batch_get_id", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.rooms.service.conf, req)
	return result, err
}

func (rooms *RoomService) BatchGetId(ctx *core.Context, optFns ...request.OptFn) *RoomBatchGetIdReqCall {
	return &RoomBatchGetIdReqCall{
		ctx:         ctx,
		rooms:       rooms,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type RoomCreateReqCall struct {
	ctx    *core.Context
	rooms  *RoomService
	body   *RoomCreateReqBody
	optFns []request.OptFn
}

func (rc *RoomCreateReqCall) Do() (*RoomCreateResult, error) {
	var result = &RoomCreateResult{}
	req := request.NewRequest("meeting_room/room/create", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.rooms.service.conf, req)
	return result, err
}

func (rooms *RoomService) Create(ctx *core.Context, body *RoomCreateReqBody, optFns ...request.OptFn) *RoomCreateReqCall {
	return &RoomCreateReqCall{
		ctx:    ctx,
		rooms:  rooms,
		body:   body,
		optFns: optFns,
	}
}

type BuildingCreateReqCall struct {
	ctx       *core.Context
	buildings *BuildingService
	body      *BuildingCreateReqBody
	optFns    []request.OptFn
}

func (rc *BuildingCreateReqCall) Do() (*BuildingCreateResult, error) {
	var result = &BuildingCreateResult{}
	req := request.NewRequest("meeting_room/building/create", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.buildings.service.conf, req)
	return result, err
}

func (buildings *BuildingService) Create(ctx *core.Context, body *BuildingCreateReqBody, optFns ...request.OptFn) *BuildingCreateReqCall {
	return &BuildingCreateReqCall{
		ctx:       ctx,
		buildings: buildings,
		body:      body,
		optFns:    optFns,
	}
}

type BuildingDeleteReqCall struct {
	ctx       *core.Context
	buildings *BuildingService
	body      *BuildingDeleteReqBody
	optFns    []request.OptFn
}

func (rc *BuildingDeleteReqCall) Do() (*response.NoData, error) {
	var result = &response.NoData{}
	req := request.NewRequest("meeting_room/building/delete", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.buildings.service.conf, req)
	return result, err
}

func (buildings *BuildingService) Delete(ctx *core.Context, body *BuildingDeleteReqBody, optFns ...request.OptFn) *BuildingDeleteReqCall {
	return &BuildingDeleteReqCall{
		ctx:       ctx,
		buildings: buildings,
		body:      body,
		optFns:    optFns,
	}
}

type RoomDeleteReqCall struct {
	ctx    *core.Context
	rooms  *RoomService
	body   *RoomDeleteReqBody
	optFns []request.OptFn
}

func (rc *RoomDeleteReqCall) Do() (*response.NoData, error) {
	var result = &response.NoData{}
	req := request.NewRequest("meeting_room/room/delete", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.rooms.service.conf, req)
	return result, err
}

func (rooms *RoomService) Delete(ctx *core.Context, body *RoomDeleteReqBody, optFns ...request.OptFn) *RoomDeleteReqCall {
	return &RoomDeleteReqCall{
		ctx:    ctx,
		rooms:  rooms,
		body:   body,
		optFns: optFns,
	}
}

type RoomListReqCall struct {
	ctx         *core.Context
	rooms       *RoomService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *RoomListReqCall) SetBuildingId(buildingId string) {
	rc.queryParams["building_id"] = buildingId
}
func (rc *RoomListReqCall) SetOrderBy(orderBy string) {
	rc.queryParams["order_by"] = orderBy
}
func (rc *RoomListReqCall) SetFields(fields string) {
	rc.queryParams["fields"] = fields
}
func (rc *RoomListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *RoomListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *RoomListReqCall) Do() (*RoomListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &RoomListResult{}
	req := request.NewRequest("meeting_room/room/list", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.rooms.service.conf, req)
	return result, err
}

func (rooms *RoomService) List(ctx *core.Context, optFns ...request.OptFn) *RoomListReqCall {
	return &RoomListReqCall{
		ctx:         ctx,
		rooms:       rooms,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type BuildingListReqCall struct {
	ctx         *core.Context
	buildings   *BuildingService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *BuildingListReqCall) SetOrderBy(orderBy string) {
	rc.queryParams["order_by"] = orderBy
}
func (rc *BuildingListReqCall) SetFields(fields string) {
	rc.queryParams["fields"] = fields
}
func (rc *BuildingListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *BuildingListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *BuildingListReqCall) Do() (*BuildingListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &BuildingListResult{}
	req := request.NewRequest("meeting_room/building/list", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.buildings.service.conf, req)
	return result, err
}

func (buildings *BuildingService) List(ctx *core.Context, optFns ...request.OptFn) *BuildingListReqCall {
	return &BuildingListReqCall{
		ctx:         ctx,
		buildings:   buildings,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type DistrictListReqCall struct {
	ctx         *core.Context
	districts   *DistrictService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *DistrictListReqCall) SetCountryId(countryId int) {
	rc.queryParams["country_id"] = countryId
}

func (rc *DistrictListReqCall) Do() (*DistrictListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &DistrictListResult{}
	req := request.NewRequest("meeting_room/district/list", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.districts.service.conf, req)
	return result, err
}

func (districts *DistrictService) List(ctx *core.Context, optFns ...request.OptFn) *DistrictListReqCall {
	return &DistrictListReqCall{
		ctx:         ctx,
		districts:   districts,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type CountryListReqCall struct {
	ctx      *core.Context
	countrys *CountryService
	optFns   []request.OptFn
}

func (rc *CountryListReqCall) Do() (*CountryListResult, error) {
	var result = &CountryListResult{}
	req := request.NewRequest("meeting_room/country/list", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.countrys.service.conf, req)
	return result, err
}

func (countrys *CountryService) List(ctx *core.Context, optFns ...request.OptFn) *CountryListReqCall {
	return &CountryListReqCall{
		ctx:      ctx,
		countrys: countrys,
		optFns:   optFns,
	}
}

type InstanceReplyReqCall struct {
	ctx       *core.Context
	instances *InstanceService
	body      *InstanceReplyReqBody
	optFns    []request.OptFn
}

func (rc *InstanceReplyReqCall) Do() (*response.NoData, error) {
	var result = &response.NoData{}
	req := request.NewRequest("meeting_room/instance/reply", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.instances.service.conf, req)
	return result, err
}

func (instances *InstanceService) Reply(ctx *core.Context, body *InstanceReplyReqBody, optFns ...request.OptFn) *InstanceReplyReqCall {
	return &InstanceReplyReqCall{
		ctx:       ctx,
		instances: instances,
		body:      body,
		optFns:    optFns,
	}
}

type BuildingUpdateReqCall struct {
	ctx       *core.Context
	buildings *BuildingService
	body      *BuildingUpdateReqBody
	optFns    []request.OptFn
}

func (rc *BuildingUpdateReqCall) Do() (*response.NoData, error) {
	var result = &response.NoData{}
	req := request.NewRequest("meeting_room/building/update", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.buildings.service.conf, req)
	return result, err
}

func (buildings *BuildingService) Update(ctx *core.Context, body *BuildingUpdateReqBody, optFns ...request.OptFn) *BuildingUpdateReqCall {
	return &BuildingUpdateReqCall{
		ctx:       ctx,
		buildings: buildings,
		body:      body,
		optFns:    optFns,
	}
}

type RoomUpdateReqCall struct {
	ctx    *core.Context
	rooms  *RoomService
	body   *RoomUpdateReqBody
	optFns []request.OptFn
}

func (rc *RoomUpdateReqCall) Do() (*response.NoData, error) {
	var result = &response.NoData{}
	req := request.NewRequest("meeting_room/room/update", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.rooms.service.conf, req)
	return result, err
}

func (rooms *RoomService) Update(ctx *core.Context, body *RoomUpdateReqBody, optFns ...request.OptFn) *RoomUpdateReqCall {
	return &RoomUpdateReqCall{
		ctx:    ctx,
		rooms:  rooms,
		body:   body,
		optFns: optFns,
	}
}