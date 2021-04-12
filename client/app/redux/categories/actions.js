import {
    CREATE_CATEGORY_REQUEST,
    DELETE_CATEGORY_REQUEST,
    LIST_CATEGORIES_REQUEST,
    UPDATE_CATEGORY_REQUEST
} from './constants';

export const listCategoriesRequest = () => ({
    type: LIST_CATEGORIES_REQUEST,
});

export const createCategoryRequest = (name) => ({
    type: CREATE_CATEGORY_REQUEST,
    name: name,
});

export const updateCategoryRequest = (id, name) => ({
    type: UPDATE_CATEGORY_REQUEST,
    id: id,
    name: name,
});

export const deleteCategoryRequest = (id) => ({
    type: DELETE_CATEGORY_REQUEST,
    id: id,
});
