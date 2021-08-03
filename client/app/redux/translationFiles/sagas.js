import {all, put, call, takeLatest} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';

import {getHeaders} from '../utils';
import {
    LIST_TRANSLATIONS_FILES_REQUEST,
    LIST_TRANSLATIONS_FILES_SUCCESS,
    LIST_TRANSLATIONS_FILES_ERROR,
    CREATE_TRANSLATION_FILES_REQUEST,
    CREATE_TRANSLATION_FILES_ERROR,
    DELETE_TRANSLATION_FILES_REQUEST,
    DELETE_TRANSLATION_FILES_ERROR,
} from './constants';

const getTranslationFilesUrl = `${apiServerURL}/api/internal/translation-files`;
const createTranslateFileUrl = `${apiServerURL}/api/internal/translation-files`;
const deleteTranslationFileUrl = `${apiServerURL}/api/internal/translation-files/`;//id

function getTranslationFilesApi(localizationId = 0, platform = '') {
    let url = getTranslationFilesUrl;
    if (localizationId > 0) {
        url = url + '?localization_id=' + localizationId;
    } else if (platform !== '') {
        url = url + '?platform=' + platform;
    } else if (localizationId > 0 && platform !== '') {
        url = url + '?localization_id=' + localizationId + '&platform=' + platform;
    }

    return fetch(url, {
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

function* getTranslationFilesWorker(action) {
    try {
        const {localizationId, platform} = action;
        const resp = yield call(getTranslationFilesApi, localizationId, platform);
        yield put({type: LIST_TRANSLATIONS_FILES_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: LIST_TRANSLATIONS_FILES_ERROR, error});
    }
}

function createTranslationFileApi(text, localizationId, platforms = [], storageType = '') {
    return fetch(createTranslateFileUrl, {
        method: 'POST',
        headers: getHeaders(),
        body: JSON.stringify({
            text, 'localization_id': localizationId, 'platforms': platforms, 'storage_type': storageType,
        }),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* createTranslationFileWorker(action) {
    try {
        const {localizationId, platforms, storageType} = action;
        yield call(createTranslationFileApi, localizationId, platforms, storageType);
        yield put({type: LIST_TRANSLATIONS_FILES_REQUEST});
    } catch (error) {
        yield put({type: CREATE_TRANSLATION_FILES_ERROR, error});
    }
}

function deleteTranslationFileApi(id) {
    return fetch(deleteTranslationFileUrl + id, {
        method: 'DELETE',
        headers: getHeaders(),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* deleteTranslationFileWorker(action) {
    try {
        const {id} = action;
        yield call(deleteTranslationFileApi, id);
        yield put({type: LIST_TRANSLATIONS_FILES_REQUEST});
    } catch (error) {
        yield put({type: DELETE_TRANSLATION_FILES_ERROR, error});
    }
}

export default function* rootSaga() {
    yield all([
        takeLatest(LIST_TRANSLATIONS_FILES_REQUEST, getTranslationFilesWorker),
        takeLatest(CREATE_TRANSLATION_FILES_REQUEST, createTranslationFileWorker),
        takeLatest(DELETE_TRANSLATION_FILES_REQUEST, deleteTranslationFileWorker),
    ]);
}
