import {
    CREATE_CATEGORY_ERROR,
    CREATE_CATEGORY_SUCCESS, DELETE_CATEGORY_ERROR, DELETE_CATEGORY_SUCCESS,
    LIST_CATEGORIES_ERROR,
    LIST_CATEGORIES_REQUEST,
    LIST_CATEGORIES_SUCCESS
} from './constants';

const initialState = {
    error: '',
    isLoading: false,
    successful: false,
    listCategories: []
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case LIST_CATEGORIES_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case LIST_CATEGORIES_SUCCESS:
            return {
                error: '',
                isLoading: false,
                successful: true,
                listCategories: action.payload,
            };
        case LIST_CATEGORIES_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case CREATE_CATEGORY_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case CREATE_CATEGORY_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case DELETE_CATEGORY_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case DELETE_CATEGORY_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        default:
            return state;
    }
}
