import {all} from 'redux-saga/effects';
import login from './login/sagas';
import profile from './profile/sagas';
import localizations from './localizations/sagas';
import identifiers from './identifiers/sagas';
import categories from './categories/sagas';
import translations from './translations/sagas';
import users from './users/sagas';
import stat from './stat/sagas';
import translationFiles from './translationFiles/sagas';
import languages from './languages/sagas';
import settings from './settings/sagas';

export default function* rootSaga() {
    yield all([
        login(),
        profile(),
        localizations(),
        identifiers(),
        categories(),
        translations(),
        users(),
        stat(),
        translationFiles(),
        languages(),
        settings()
    ]);
}
