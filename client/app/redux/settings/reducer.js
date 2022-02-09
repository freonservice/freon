import {
    LIST_SETTINGS_ERROR,
    LIST_SETTINGS_REQUEST,
    LIST_SETTINGS_SUCCESS,
    UPDATE_SETTING_STORAGE_ERROR,
    UPDATE_SETTING_STORAGE_SUCCESS,
    UPDATE_SETTING_TRANSLATION_ERROR,
    UPDATE_SETTING_TRANSLATION_SUCCESS
} from './constants';

const initialState = {
    error: '',
    isLoading: false,
    successful: false,
    settings: {}
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case LIST_SETTINGS_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case LIST_SETTINGS_SUCCESS:
            return {
                error: '',
                isLoading: false,
                successful: true,
                settings: action.payload,
            };
        case LIST_SETTINGS_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case UPDATE_SETTING_STORAGE_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case UPDATE_SETTING_STORAGE_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case UPDATE_SETTING_TRANSLATION_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case UPDATE_SETTING_TRANSLATION_ERROR:
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
