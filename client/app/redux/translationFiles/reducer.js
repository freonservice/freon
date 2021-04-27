import {
    CREATE_TRANSLATION_FILES_ERROR,
    CREATE_TRANSLATION_FILES_SUCCESS, DELETE_TRANSLATION_FILES_ERROR, DELETE_TRANSLATION_FILES_REQUEST,
    LIST_TRANSLATIONS_FILES_ERROR,
    LIST_TRANSLATIONS_FILES_REQUEST,
    LIST_TRANSLATIONS_FILES_SUCCESS
} from './constants';

const initialState = {
    error: '',
    isLoading: false,
    successful: false,
    listTranslationFiles: []
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case LIST_TRANSLATIONS_FILES_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case LIST_TRANSLATIONS_FILES_SUCCESS:
            return {
                error: '',
                isLoading: false,
                successful: true,
                listTranslationFiles: action.payload,
            };
        case LIST_TRANSLATIONS_FILES_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case CREATE_TRANSLATION_FILES_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case CREATE_TRANSLATION_FILES_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case DELETE_TRANSLATION_FILES_REQUEST:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case DELETE_TRANSLATION_FILES_ERROR:
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
