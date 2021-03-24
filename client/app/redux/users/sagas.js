import {all, put, call, takeLatest} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';

import {getHeaders} from '../utils';
import {
    GET_USER_LIST_SUCCESS,
    GET_USER_LIST_ERROR,
    GET_USER_LIST_REQUEST, CREATE_USER_REQUEST,
} from './constants';

const getUserListUrl = `${apiServerURL}/api/internal/users`;
const createUserUrl = `${apiServerURL}/api/internal/user/register`;

function getUserListApi() {
    return fetch(getUserListUrl, {
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

function* getUserListWorker() {
    try {
        const resp = yield call(getUserListApi);
        yield put({type: GET_USER_LIST_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: GET_USER_LIST_ERROR, error});
    }
}

function createUserApi(action) {
    const {email, first_name, second_name, password, repeat_password, role} = action;
    return fetch(createUserUrl, {
        method: 'POST',
        headers: getHeaders(),
        body: JSON.stringify({email, first_name, second_name, password, repeat_password, role}),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* createUserWorker(action) {
    try {
        yield call(createUserApi, action);
        yield put({type: GET_USER_LIST_REQUEST});
    } catch (error) {
        // yield put({type: CREATE_TRANSLATION_ERROR, error});
    }
}

export default function* rootSaga() {
    yield all([
        takeLatest(GET_USER_LIST_REQUEST, getUserListWorker),
        takeLatest(CREATE_USER_REQUEST, createUserWorker),
    ]);
}
