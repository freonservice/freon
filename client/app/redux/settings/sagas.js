import {all, put, call, takeLatest} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';

import {getHeaders} from '../utils';
import {
    LIST_SETTINGS_SUCCESS,
    LIST_SETTINGS_ERROR,
    LIST_SETTINGS_REQUEST,
    UPDATE_SETTING_STORAGE_SUCCESS,
    UPDATE_SETTING_STORAGE_ERROR,
    UPDATE_SETTING_STORAGE_REQUEST,
    UPDATE_SETTING_TRANSLATION_REQUEST, UPDATE_SETTING_TRANSLATION_SUCCESS, UPDATE_SETTING_TRANSLATION_ERROR,
} from './constants';

const getSettingsUrl = `${apiServerURL}/api/internal/settings`;
const updateSettingStorageUrl = `${apiServerURL}/api/internal/setting/storage`;
const updateSettingTranslationUrl = `${apiServerURL}/api/internal/setting/translation`;

function getSettingsApi() {
    return fetch(getSettingsUrl, {
        method: 'GET',
        headers: getHeaders(),
    })
        .then(handleApiErrors)
        .then(response => response.json())
        .then(json => json)
        .catch((error) => {
            throw error;
        });
}

function* getSettingsWorker() {
    try {
        const resp = yield call(getSettingsApi);
        yield put({type: LIST_SETTINGS_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: LIST_SETTINGS_ERROR, error});
    }
}

function updateSettingStorageApi(action) {
    const {use} = action;
    return fetch(updateSettingStorageUrl, {
        method: 'PUT',
        headers: getHeaders(),
        body: JSON.stringify({
            use,
        }),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* updateSettingStorageWorker(action) {
    try {
        yield call(updateSettingStorageApi, action);
        yield put({type: UPDATE_SETTING_STORAGE_SUCCESS});
    } catch (error) {
        yield put({type: UPDATE_SETTING_STORAGE_ERROR, error});
    }
}

function updateSettingTranslationApi(action) {
    const {use, auto, main_language} = action;
    return fetch(updateSettingStorageUrl, {
        method: 'PUT',
        headers: getHeaders(),
        body: JSON.stringify({
            use, auto, main_language
        }),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* updateSettingTranslationWorker(action) {
    try {
        yield call(updateSettingTranslationApi, action);
        yield put({type: UPDATE_SETTING_TRANSLATION_SUCCESS});
    } catch (error) {
        yield put({type: UPDATE_SETTING_TRANSLATION_ERROR, error});
    }
}

export default function* rootSaga() {
    yield all([
        takeLatest(LIST_SETTINGS_REQUEST, getSettingsWorker),
        takeLatest(UPDATE_SETTING_STORAGE_REQUEST, updateSettingStorageWorker),
        takeLatest(UPDATE_SETTING_TRANSLATION_REQUEST, updateSettingTranslationWorker),
    ]);
}
