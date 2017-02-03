package api

/************************************************
  generated by IDE. for [PacketFilterAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *PacketFilterAPI) Reset() *PacketFilterAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *PacketFilterAPI) Offset(offset int) *PacketFilterAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *PacketFilterAPI) Limit(limit int) *PacketFilterAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *PacketFilterAPI) Include(key string) *PacketFilterAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *PacketFilterAPI) Exclude(key string) *PacketFilterAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルター
func (api *PacketFilterAPI) FilterBy(key string, value interface{}) *PacketFilterAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *PacketFilterAPI) FilterMultiBy(key string, value interface{}) *PacketFilterAPI {
	api.filterBy(key, value, true)
	return api
}

// WithNameLike 名称条件
func (api *PacketFilterAPI) WithNameLike(name string) *PacketFilterAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *PacketFilterAPI) WithTag(tag string) *PacketFilterAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *PacketFilterAPI) WithTags(tags []string) *PacketFilterAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *PacketFilterAPI) WithSizeGib(size int) *PacketFilterAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *PacketFilterAPI) WithSharedScope() *PacketFilterAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *PacketFilterAPI) WithUserScope() *PacketFilterAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

// SortBy 指定キーでのソート
func (api *PacketFilterAPI) SortBy(key string, reverse bool) *PacketFilterAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *PacketFilterAPI) SortByName(reverse bool) *PacketFilterAPI {
	api.sortByName(reverse)
	return api
}

// func (api *PacketFilterAPI) SortBySize(reverse bool) *PacketFilterAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// New 新規作成用パラメーター作成
func (api *PacketFilterAPI) New() *sacloud.PacketFilter {
	return sacloud.CreateNewPacketFilter()
}

// Create 新規作成
func (api *PacketFilterAPI) Create(value *sacloud.PacketFilter) (*sacloud.PacketFilter, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

// Read 読み取り
func (api *PacketFilterAPI) Read(id int64) (*sacloud.PacketFilter, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

// Update 更新
func (api *PacketFilterAPI) Update(id int64, value *sacloud.PacketFilter) (*sacloud.PacketFilter, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

// Delete 削除
func (api *PacketFilterAPI) Delete(id int64) (*sacloud.PacketFilter, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

/************************************************
  Inner functions
************************************************/

func (api *PacketFilterAPI) setStateValue(setFunc func(*sacloud.Request)) *PacketFilterAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *PacketFilterAPI) request(f func(*sacloud.Response) error) (*sacloud.PacketFilter, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.PacketFilter, nil
}

func (api *PacketFilterAPI) createRequest(value *sacloud.PacketFilter) *sacloud.Request {
	req := &sacloud.Request{}
	req.PacketFilter = value
	return req
}
