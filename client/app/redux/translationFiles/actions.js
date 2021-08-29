import {
    LIST_TRANSLATIONS_FILES_REQUEST,
    CREATE_TRANSLATION_FILES_REQUEST,
    DELETE_TRANSLATION_FILES_REQUEST,
} from './constants';

export const listTranslationFilesRequest = (localizationId = 0, platformType = 0) => ({
    type: LIST_TRANSLATIONS_FILES_REQUEST,
    localizationId: localizationId,
    platformType: platformType,
});

export const createTranslationFilesRequest = (localizationId, storageType, platforms = []) => ({
    type: CREATE_TRANSLATION_FILES_REQUEST,
    localizationId: localizationId,
    storageType: storageType,
    platforms: platforms,
});

export const deleteTranslationFileRequest = (id) => ({
    type: DELETE_TRANSLATION_FILES_REQUEST,
    id: id,
});
