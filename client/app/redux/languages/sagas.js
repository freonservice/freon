import {all, put, call, takeLatest} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';

import {getHeaders} from '../utils';
import {
    LIST_LANGUAGES_REQUEST,
    LIST_LANGUAGES_ERROR,
    LIST_LANGUAGES_SUCCESS,
} from './constants';

const getSupportedLanguagesUrl = `${apiServerURL}/api/internal/supported-languages`;

function getSupportedLanguagesApi() {
    return fetch(getSupportedLanguagesUrl, {
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

function* getSupportedLanguagesWorker() {
    try {
        const resp = yield call(getSupportedLanguagesApi);
        yield put({type: LIST_LANGUAGES_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: LIST_LANGUAGES_ERROR, error});
    }
}

export default function* rootSaga() {
    yield all([
        takeLatest(LIST_LANGUAGES_REQUEST, getSupportedLanguagesWorker),
    ]);
}
