import {
    CREATE_TRANSLATION_REQUEST,
    LIST_TRANSLATIONS_REQUEST,
    UPDATE_TRANSLATION_REQUEST,
    HIDE_TRANSLATION_REQUEST
} from './constants';

export const listTranslationsRequest = (localizationId = 0) => ({
    type: LIST_TRANSLATIONS_REQUEST,
    localizationId: localizationId,
});

export const createTranslationRequest = (text, localizationId, identifierId) => ({
    type: CREATE_TRANSLATION_REQUEST,
    text: text,
    localizationId: localizationId,
    identifierId: identifierId,
});

export const updateTranslationRequest = (id, text) => ({
    type: UPDATE_TRANSLATION_REQUEST,
    id: id,
    text: text,
});

export const hideTranslationRequest = (id, hide = true) => ({
    type: HIDE_TRANSLATION_REQUEST,
    id: id,
    hide: hide
});
