import {
    CREATE_TRANSLATION_REQUEST,
    LIST_TRANSLATIONS_REQUEST,
    UPDATE_TRANSLATION_REQUEST,
    UPDATE_STATUS_TRANSLATION_REQUEST
} from './constants';

export const listTranslationsRequest = (localizationId = 0) => ({
    type: LIST_TRANSLATIONS_REQUEST,
    localizationId: localizationId,
});

export const createTranslationRequest = (singular, plural, localizationId, identifierId) => ({
    type: CREATE_TRANSLATION_REQUEST,
    singular: singular,
    plural: plural,
    localizationId: localizationId,
    identifierId: identifierId,
});

export const updateTranslationRequest = (id, singular, plural) => ({
    type: UPDATE_TRANSLATION_REQUEST,
    id: id,
    singular: singular,
    plural: plural,
});

export const updateStatusTranslationRequest = (id, status = 0) => ({
    type: UPDATE_STATUS_TRANSLATION_REQUEST,
    id: id,
    status: status
});
