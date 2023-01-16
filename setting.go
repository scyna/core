package scyna

import (
	"encoding/json"
	"strconv"
	"sync"

	scyna_proto "github.com/scyna/core/proto/generated"
)

type settings struct {
	data  map[string]string /*cache*/
	mutex sync.Mutex
}

func (s *settings) Remove(key string) bool {
	request := scyna_proto.RemoveSettingRequest{Context: context, Key: key}
	var response scyna_proto.Error
	if err := callEndpoint(SETTING_REMOVE_URL, &request, &response); err.Code() == OK.Code() {
		s.removed(key)
		return true
	}
	return false
}

func (s *settings) Write(key string, value string) bool {
	request := scyna_proto.WriteSettingRequest{Context: context, Key: key, Value: value}
	var response scyna_proto.Error
	if err := callEndpoint(SETTING_WRITE_URL, &request, &response); err.Code() == OK.Code() {
		s.updated(key, value)
		return true
	}
	return false
}

func (s *settings) ReadString(key string) (bool, string) {
	/*from cache*/
	s.mutex.Lock()
	if val, ok := s.data[key]; ok {
		s.mutex.Unlock()
		return true, val
	}
	s.mutex.Unlock()

	/*from manager*/
	request := scyna_proto.ReadSettingRequest{Context: context, Key: key}
	var response scyna_proto.ReadSettingResponse
	if err := callEndpoint(SETTING_READ_URL, &request, &response); err.Code() == OK.Code() {
		s.updated(key, response.Value)
		return true, response.Value
	}
	return false, ""
}

func (s *settings) ReadInt(key string) (bool, int) {
	if ok, val := s.ReadString(key); ok {
		if i, err := strconv.Atoi(val); err != nil {
			return false, 0
		} else {
			return true, i
		}
	}
	return false, 0
}

func (s *settings) ReadBool(key string) (bool, bool) {
	/*TODO*/
	return false, false
}

func (s *settings) ReadObject(key string, value interface{}) bool {
	if ok, val := s.ReadString(key); ok {
		if err := json.Unmarshal([]byte(val), value); err != nil {
			LOG.Warning("ReadObjectSetting: " + err.Error())
			return false
		}
		return true
	}
	return false
}

func UpdateSettingHandler(data *scyna_proto.SettingUpdatedSignal) {
	if data.Context == context {
		Settings.updated(data.Key, data.Value)
	}
}

func RemoveSettingHandler(data *scyna_proto.SettingRemovedSignal) {
	if data.Context == context {
		Settings.removed(data.Key)
	}
}

func (s *settings) updated(key string, value string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[key] = value
}

func (s *settings) removed(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, key)
}

func (s *settings) Init() {
	s.data = make(map[string]string)
}
