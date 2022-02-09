import {
    LIST_LANGUAGES_ERROR,
    LIST_LANGUAGES_REQUEST,
    LIST_LANGUAGES_SUCCESS,
} from './constants';

const initialState = {
    error: '',
    isLoading: false,
    successful: false,
    listLanguages: []
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case LIST_LANGUAGES_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case LIST_LANGUAGES_SUCCESS:
            return {
                error: '',
                isLoading: false,
                successful: true,
                listLanguages: action.payload,
            };
        case LIST_LANGUAGES_ERROR:
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
