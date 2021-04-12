import {
    LIST_IDENTIFIERS_ERROR,
    LIST_IDENTIFIERS_REQUEST,
    LIST_IDENTIFIERS_SUCCESS,
    CREATE_IDENTIFIER_ERROR,
    CREATE_IDENTIFIER_SUCCESS,
    UPDATE_IDENTIFIER_ERROR,
    DELETE_IDENTIFIER_ERROR,
    DELETE_IDENTIFIER_SUCCESS,
} from './constants';

const initialState = {
    error: '',
    isLoading: false,
    successful: false,
    listIdentifiers: []
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case LIST_IDENTIFIERS_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case LIST_IDENTIFIERS_SUCCESS:
            return {
                error: '',
                isLoading: false,
                successful: true,
                listIdentifiers: action.payload,
            };
        case LIST_IDENTIFIERS_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case CREATE_IDENTIFIER_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case CREATE_IDENTIFIER_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case UPDATE_IDENTIFIER_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case DELETE_IDENTIFIER_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case DELETE_IDENTIFIER_ERROR:
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
