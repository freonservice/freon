import {all, put, call, takeLatest} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';

import {
    LIST_LOCALIZATIONS_REQUEST,
    LIST_LOCALIZATIONS_ERROR,
    LIST_LOCALIZATIONS_SUCCESS,
    CREATE_LOCALIZATION_ERROR,
    CREATE_LOCALIZATION_REQUEST
} from './constants';

import {getHeaders} from '../utils';

const getLocalizationsUrl = `${apiServerURL}/api/internal/localizations`;
const createLocalizationUrl = `${apiServerURL}/api/internal/localization`;

function getLocalizationsApi() {
    return fetch(getLocalizationsUrl, {
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

function* getLocalizationsWorker() {
    try {
        const resp = yield call(getLocalizationsApi);
        yield put({type: LIST_LOCALIZATIONS_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: LIST_LOCALIZATIONS_ERROR, error});
    }
}

function createLocalizationApi(locale, langName) {
    return fetch(createLocalizationUrl, {
        method: 'POST',
        headers: getHeaders(),
        body: JSON.stringify({'locale': locale, 'lang_name': langName}),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* createLocalizationWorker(action) {
    try {
        const {locale, langName} = action;
        yield call(createLocalizationApi, locale, langName);
        yield put({type: LIST_LOCALIZATIONS_REQUEST});
    } catch (error) {
        yield put({type: CREATE_LOCALIZATION_ERROR, error});
    }
}

export default function* rootSaga() {
    yield all([
        takeLatest(LIST_LOCALIZATIONS_REQUEST, getLocalizationsWorker),
        takeLatest(CREATE_LOCALIZATION_REQUEST, createLocalizationWorker)
    ]);
}
