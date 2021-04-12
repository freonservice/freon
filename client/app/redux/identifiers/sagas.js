import {all, put, call, takeLatest} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';

import {getHeaders} from '../utils';
import {
    LIST_IDENTIFIERS_ERROR,
    LIST_IDENTIFIERS_REQUEST,
    LIST_IDENTIFIERS_SUCCESS,
    CREATE_IDENTIFIER_REQUEST,
    CREATE_IDENTIFIER_ERROR,
    UPDATE_IDENTIFIER_ERROR,
    UPDATE_IDENTIFIER_REQUEST,
    DELETE_IDENTIFIER_ERROR,
    DELETE_IDENTIFIER_REQUEST,
} from './constants';

const getIdentifiersUrl = `${apiServerURL}/api/internal/identifiers`;
const createIdentifierUrl = `${apiServerURL}/api/internal/identifier`;
const deleteOrUpdateIdentifierUrl = `${apiServerURL}/api/internal/identifier/`;

function getIdentifiersApi() {
    return fetch(getIdentifiersUrl, {
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

function* getIdentifiersWorker() {
    try {
        const resp = yield call(getIdentifiersApi);
        yield put({type: LIST_IDENTIFIERS_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: LIST_IDENTIFIERS_ERROR, error});
    }
}

function createIdentifierApi(action) {
    const {name, description, example_text, categoryId, platforms, namedList} = action;
    return fetch(createIdentifierUrl, {
        method: 'POST',
        headers: getHeaders(),
        body: JSON.stringify({
            name,
            description,
            example_text,
            'category_id': categoryId,
            platforms,
            'named_list': namedList
        }),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* createIdentifierWorker(action) {
    try {
        yield call(createIdentifierApi, action);
        yield put({type: LIST_IDENTIFIERS_REQUEST});
    } catch (error) {
        yield put({type: CREATE_IDENTIFIER_ERROR, error});
    }
}

function deleteIdentifierApi(id) {
    return fetch(deleteOrUpdateIdentifierUrl + id, {
        method: 'DELETE',
        headers: getHeaders(),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* deleteIdentifierWorker(action) {
    try {
        const {id} = action;
        yield call(deleteIdentifierApi, id);
        yield put({type: LIST_IDENTIFIERS_REQUEST});
    } catch (error) {
        yield put({type: DELETE_IDENTIFIER_ERROR, error});
    }
}

function updateIdentifierApi(action) {
    const {id, name, description, example_text, categoryId, platforms, namedList} = action;
    return fetch(deleteOrUpdateIdentifierUrl + id, {
        method: 'PUT',
        headers: getHeaders(),
        body: JSON.stringify({
            name,
            description,
            example_text,
            'category_id': categoryId,
            platforms,
            'named_list': namedList
        }),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* updateIdentifierWorker(action) {
    try {
        yield call(updateIdentifierApi, action);
        yield put({type: LIST_IDENTIFIERS_REQUEST});
    } catch (error) {
        yield put({type: UPDATE_IDENTIFIER_ERROR, error});
    }
}

export default function* rootSaga() {
    yield all([
        takeLatest(LIST_IDENTIFIERS_REQUEST, getIdentifiersWorker),
        takeLatest(CREATE_IDENTIFIER_REQUEST, createIdentifierWorker),
        takeLatest(UPDATE_IDENTIFIER_REQUEST, updateIdentifierWorker),
        takeLatest(DELETE_IDENTIFIER_REQUEST, deleteIdentifierWorker)
    ]);
}
