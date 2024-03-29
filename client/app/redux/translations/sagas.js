import {all, put, call, takeLatest} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';

import {getHeaders} from '../utils';
import {
    LIST_TRANSLATIONS_ERROR,
    LIST_TRANSLATIONS_REQUEST,
    LIST_TRANSLATIONS_SUCCESS,
    CREATE_TRANSLATION_ERROR,
    CREATE_TRANSLATION_REQUEST,
    UPDATE_TRANSLATION_ERROR,
    UPDATE_TRANSLATION_REQUEST,
    UPDATE_STATUS_TRANSLATION_ERROR,
    UPDATE_STATUS_TRANSLATION_REQUEST,
} from './constants';

const getTranslationsUrl = `${apiServerURL}/api/internal/translations`;
const createTranslateUrl = `${apiServerURL}/api/internal/translation`;
const updateTranslationUrl = `${apiServerURL}/api/internal/translation/`;
// const hideTranslationUrl = `${apiServerURL}/api/internal/translation/hide/`;

function getTranslationsApi(localizationId = 0) {
    let url = getTranslationsUrl;
    if (localizationId > 0) {
        url = url + '?localization_id=' + localizationId;
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

function* getTranslationsWorker(action) {
    try {
        const {localizationId} = action;
        const resp = yield call(getTranslationsApi, localizationId);
        yield put({type: LIST_TRANSLATIONS_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: LIST_TRANSLATIONS_ERROR, error});
    }
}

function createTranslationApi(text, localizationId, identifierId) {
    return fetch(createTranslateUrl, {
        method: 'POST',
        headers: getHeaders(),
        body: JSON.stringify({text, 'localization_id': localizationId, 'identifier_id': identifierId}),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* createTranslationWorker(action) {
    try {
        const {text, localizationId, identifierId} = action;
        yield call(createTranslationApi, text, localizationId, identifierId);
        yield put({type: LIST_TRANSLATIONS_REQUEST});
    } catch (error) {
        yield put({type: CREATE_TRANSLATION_ERROR, error});
    }
}

function statusTranslationApi(id, status) {
    let apiUrl = `${apiServerURL}/api/internal/translation/` + id + '/status/' + status;
    return fetch(apiUrl, {
        method: 'PUT',
        headers: getHeaders(),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* statusTranslationWorker(action) {
    try {
        const {id, status} = action;
        yield call(statusTranslationApi, id, status);
        yield put({type: LIST_TRANSLATIONS_REQUEST});
    } catch (error) {
        yield put({type: UPDATE_STATUS_TRANSLATION_ERROR, error});
    }
}

function updateTranslationApi(id, singular, plural) {
    return fetch(updateTranslationUrl + id, {
        method: 'PUT',
        headers: getHeaders(),
        body: JSON.stringify({singular, plural}),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* updateTranslationWorker(action) {
    try {
        const {id, singular, plural} = action;
        yield call(updateTranslationApi, id, singular, plural);
        yield put({type: LIST_TRANSLATIONS_REQUEST});
    } catch (error) {
        yield put({type: UPDATE_TRANSLATION_ERROR, error});
    }
}

export default function* rootSaga() {
    yield all([
        takeLatest(LIST_TRANSLATIONS_REQUEST, getTranslationsWorker),
        takeLatest(CREATE_TRANSLATION_REQUEST, createTranslationWorker),
        takeLatest(UPDATE_STATUS_TRANSLATION_REQUEST, statusTranslationWorker),
        takeLatest(UPDATE_TRANSLATION_REQUEST, updateTranslationWorker)
    ]);
}
