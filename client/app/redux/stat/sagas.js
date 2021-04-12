import {all, put, call, takeEvery} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';
import {getHeaders} from '../utils';
import {STAT_REQUEST, STAT_ERROR, STAT_SUCCESS} from './constants';

const statInternalUrl = `${apiServerURL}/api/internal/statistic`;

function statInternalApi() {
    return fetch(statInternalUrl, {
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

function* statInternalWorker() {
    try {
        const resp = yield call(statInternalApi);
        yield put({type: STAT_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: STAT_ERROR, error});
    }
}

export default function* rootSaga() {
    yield all([
        takeEvery(STAT_REQUEST, statInternalWorker),
    ]);
}
