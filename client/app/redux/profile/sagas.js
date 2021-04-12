import {all, put, call, takeLatest} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';

import {getHeaders} from '../utils';
import {
    PROFILE_REQUEST,
    PROFILE_ERROR,
    PROFILE_SUCCESS,
    UPDATE_PROFILE_REQUEST,
    UPDATE_USER_PASSWORD_REQUEST,
} from './constants';

const userMeUrl = `${apiServerURL}/api/internal/user/me`;
const userUpdatePasswordUrl = `${apiServerURL}/api/internal/user/change-password`;
const userUpdateProfileUrl = `${apiServerURL}/api/internal/user/change-profile`;

function userMeApi() {
    return fetch(userMeUrl, {
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

function* userMeWorker() {
    try {
        const resp = yield call(userMeApi);
        yield put({type: PROFILE_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: PROFILE_ERROR, error});
    }
}

function userUpdateProfileApi(action) {
    const {email, first_name, second_name, role, status, user_id} = action;
    return fetch(userUpdateProfileUrl, {
        method: 'PUT',
        headers: getHeaders(),
        body: JSON.stringify({email, first_name, second_name, role, status, user_id}),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* userUpdateProfileWorker(action) {
    try {
        // const {email, first_name, second_name, role, user_id} = action;
        yield call(userUpdateProfileApi, action);
        yield put({type: PROFILE_REQUEST});
    } catch (error) {
        // yield put({type: UPDATE_TRANSLATION_ERROR, error});
    }
}

function userUpdatePasswordApi(old_password, new_password, repeat_password) {
    return fetch(userUpdatePasswordUrl, {
        method: 'PUT',
        headers: getHeaders(),
        body: JSON.stringify({old_password, new_password, repeat_password}),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* userUpdatePasswordWorker(action) {
    try {
        const {old_password, new_password, repeat_password} = action;
        yield call(userUpdatePasswordApi, old_password, new_password, repeat_password);
        yield put({type: PROFILE_REQUEST});
    } catch (error) {
        // yield put({type: UPDATE_TRANSLATION_ERROR, error});
    }
}


export default function* rootSaga() {
    yield all([
        takeLatest(PROFILE_REQUEST, userMeWorker),
        takeLatest(UPDATE_PROFILE_REQUEST, userUpdateProfileWorker),
        takeLatest(UPDATE_USER_PASSWORD_REQUEST, userUpdatePasswordWorker),
    ]);
}
