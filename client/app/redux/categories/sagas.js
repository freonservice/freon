import {all, put, call, takeLatest} from 'redux-saga/effects';
import {handleApiErrors, apiServerURL} from '../api-errors';

import {getHeaders} from '../utils';
import {
    LIST_CATEGORIES_ERROR,
    LIST_CATEGORIES_REQUEST,
    LIST_CATEGORIES_SUCCESS,
    CREATE_CATEGORY_ERROR,
    CREATE_CATEGORY_REQUEST,
    UPDATE_CATEGORY_ERROR,
    UPDATE_CATEGORY_REQUEST,
    DELETE_CATEGORY_ERROR,
    DELETE_CATEGORY_REQUEST,
} from './constants';

const getCategoriesUrl = `${apiServerURL}/api/internal/categories`;
const createCategoryUrl = `${apiServerURL}/api/internal/category`;
const deleteOrUpdateCategoryUrl = `${apiServerURL}/api/internal/category/`;

function getCategoriesApi() {
    return fetch(getCategoriesUrl, {
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

function* getCategoriesWorker() {
    try {
        const resp = yield call(getCategoriesApi);
        yield put({type: LIST_CATEGORIES_SUCCESS, payload: resp});
    } catch (error) {
        yield put({type: LIST_CATEGORIES_ERROR, error});
    }
}

function createCategoryApi(name) {
    return fetch(createCategoryUrl, {
        method: 'POST',
        headers: getHeaders(),
        body: JSON.stringify({name}),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* createCategoryWorker(action) {
    try {
        const {name} = action;
        yield call(createCategoryApi, name);
        yield put({type: LIST_CATEGORIES_REQUEST});
    } catch (error) {
        yield put({type: CREATE_CATEGORY_ERROR, error});
    }
}

function updateCategoryApi(id, name) {
    return fetch(deleteOrUpdateCategoryUrl + id, {
        method: 'PUT',
        headers: getHeaders(),
        body: JSON.stringify({name}),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* updateCategoryWorker(action) {
    try {
        const {id, name} = action;
        yield call(updateCategoryApi, id, name);
        yield put({type: LIST_CATEGORIES_REQUEST});
    } catch (error) {
        yield put({type: UPDATE_CATEGORY_ERROR, error});
    }
}

function deleteCategoryApi(id) {
    return fetch(deleteOrUpdateCategoryUrl + id, {
        method: 'DELETE',
        headers: getHeaders(),
    })
        .then(handleApiErrors)
        .catch((error) => {
            throw error;
        });
}

function* deleteCategoryWorker(action) {
    try {
        const {id} = action;
        yield call(deleteCategoryApi, id);
        yield put({type: LIST_CATEGORIES_REQUEST});
    } catch (error) {
        yield put({type: DELETE_CATEGORY_ERROR, error});
    }
}

export default function* rootSaga() {
    yield all([
        takeLatest(LIST_CATEGORIES_REQUEST, getCategoriesWorker),
        takeLatest(CREATE_CATEGORY_REQUEST, createCategoryWorker),
        takeLatest(UPDATE_CATEGORY_REQUEST, updateCategoryWorker),
        takeLatest(DELETE_CATEGORY_REQUEST, deleteCategoryWorker),
    ]);
}
