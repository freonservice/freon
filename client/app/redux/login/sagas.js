import {all, cancelled, put, call, takeLatest} from "redux-saga/effects";
import {handleApiErrors, apiServerURL} from "../api-errors";
import {
    LOGIN_REQUEST,
    LOGIN_SUCCESS,
    LOGIN_ERROR,
    LOGOUT_COMPLETED,
    LOGOUT_REQUEST,
    LOGOUT_ERROR,
} from "./constants";

import {forwardTo, getHeaders} from "../utils";

const loginUrl = `${apiServerURL}/api/internal/login`;
const logoutUrl = `${apiServerURL}/api/internal/logout`;

function loginApi(email, password) {
    return fetch(loginUrl, {
        method: "POST",
        headers: getHeaders(),
        body: JSON.stringify({email, password}),
    })
        .then(handleApiErrors)
        .then(response => response.json())
        .then(json => json)
        .catch((error) => {
            throw error;
        });
}

function* loginWorker(action) {
    let resp;
    try {
        const {email, password} = action;
        resp = yield call(loginApi, email, password);
        yield put({type: LOGIN_SUCCESS, payload: resp});
        localStorage.setItem("token", JSON.stringify(resp.token));
        yield call(forwardTo, "/");
    } catch (error) {
        yield put({type: LOGIN_ERROR, error});
    } finally {
        if (yield cancelled()) {
            yield call(forwardTo, "/login");
        }
    }
    return resp;
}

function logoutApi() {
    return fetch(logoutUrl, {
        method: "POST",
        headers: getHeaders(),
    }).then(handleApiErrors)
        .then(response => response.json())
        .then(json => json)
        .catch((error) => {
            throw error;
        });
}

function* logoutWorker() {
    try {
        yield call(logoutApi);
        yield put({type: LOGOUT_COMPLETED});
        yield call(forwardTo, "/login");
    } catch (error) {
        yield put({type: LOGOUT_ERROR, error});
    }
}


export default function* rootSaga() {
    yield all([
        takeLatest(LOGIN_REQUEST, loginWorker),
        takeLatest(LOGOUT_REQUEST, logoutWorker)
    ]);
}
