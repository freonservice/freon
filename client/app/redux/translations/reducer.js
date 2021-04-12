import {
    LIST_TRANSLATIONS_ERROR,
    LIST_TRANSLATIONS_REQUEST,
    LIST_TRANSLATIONS_SUCCESS,
    CREATE_TRANSLATION_ERROR,
    CREATE_TRANSLATION_SUCCESS,
    UPDATE_TRANSLATION_ERROR,
    HIDE_TRANSLATION_SUCCESS,
    HIDE_TRANSLATION_ERROR,
} from './constants';

const initialState = {
    error: '',
    isLoading: false,
    successful: false,
    listTranslations: []
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case LIST_TRANSLATIONS_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case LIST_TRANSLATIONS_SUCCESS:
            return {
                error: '',
                isLoading: false,
                successful: true,
                listTranslations: action.payload,
            };
        case LIST_TRANSLATIONS_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case CREATE_TRANSLATION_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case CREATE_TRANSLATION_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case UPDATE_TRANSLATION_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case HIDE_TRANSLATION_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case HIDE_TRANSLATION_ERROR:
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
