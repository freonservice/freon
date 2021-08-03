import React from 'react';

import {Container} from './../../components';
import {HeaderMain} from '../components/HeaderMain';
import {TranslationFilesTable} from './Table';
import {ToastContainer} from 'react-toastify';
import {connect} from 'react-redux';
import {listTranslationFilesRequest} from '../../redux/translationFiles/actions';
import * as PropTypes from 'prop-types';

class TranslationFiles extends React.Component {
    constructor(props) {
        super(props);

        this.props.listTranslationFilesRequest();
    }

    render() {
        return (
            <Container>
                <HeaderMain
                    title="Translation Files"
                    className="mb-5 mt-4"
                />
                <TranslationFilesTable
                    listTranslationFiles={this.props.listTranslationFiles}
                    handleDownloadTranslationFile={null}
                    handleDeleteTranslationFile={null}
                />
                <ToastContainer
                    position="top-right"
                    autoClose={3000}
                    draggable={false}
                    hideProgressBar={true}
                    limit={3}
                />
            </Container>
        );
    }
}

TranslationFiles.propTypes = {
    listTranslationFiles: PropTypes.array,
    listTranslationFilesRequest: PropTypes.func.isRequired,
    errorMsg: PropTypes.string,
};

const mapStateToProps = (state) => ({
    listTranslationFiles: state.translationFiles.listTranslationFiles,
    errorMsg: state.translationFiles.error,
});

const mapDispatchToProps = {
    listTranslationFilesRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(TranslationFiles);