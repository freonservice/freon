import React from 'react';
import {
    ButtonGroup,
    Button, Card, CardBody, Table
} from './../../components';
import PropTypes from 'prop-types';
import moment from 'moment';


export class TranslationFilesTable extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        const {listTranslationFiles} = this.props;
        return (
            <Card className="mb-3">
                <CardBody>
                    <Table className="mb-0" bordered responsive>
                        <thead>
                        <tr>
                            <th>ID</th>
                            <th>Name</th>
                            <th>Path</th>
                            <th>Platform</th>
                            <th>Storage Type</th>
                            <th>Created At</th>
                            <th className="text-right">
                                Actions
                            </th>
                        </tr>
                        </thead>
                        <tbody>
                        <React.Fragment>
                            {
                                listTranslationFiles.map((product) => (
                                        <tr key={product.id}>
                                            <td className="align-middle">
                                            <span className="text-inverse">
                                                {product.id}
                                            </span>
                                            </td>
                                            <td className="align-middle">
                                                {product.name}
                                            </td>
                                            <td className="align-middle">
                                                {product.path}
                                            </td>
                                            <td className="align-middle">
                                                {product.platform}
                                                {/*<i className={`fa fa-fw fa-${product.platform}`} aria-hidden="true"/>*/}
                                                {/*{*/}
                                                {/*    const v = value === 'web' ? 'chrome' : value;*/}
                                                {/*    return <i className={`fa fa-fw fa-${v}`} aria-hidden="true"/>;*/}
                                                {/*}*/}
                                            </td>
                                            <td className="align-middle">
                                                {product.storage_type}
                                            </td>
                                            <td className="align-middle">
                                                {moment.unix(product.created_at).format('DD-MMM-YYYY')}                                            </td>
                                            <td className="text-right">
                                                <ButtonGroup>
                                                    <Button
                                                        color="link"
                                                        className="text-decoration-none"
                                                    >
                                                        <i className="fa fa-download"/>
                                                    </Button>
                                                    <Button
                                                        color="link"
                                                        className="text-decoration-none"
                                                    >
                                                        <i className="fa fa-close"/>
                                                    </Button>
                                                </ButtonGroup>
                                            </td>
                                        </tr>
                                    )
                                )
                            }
                        </React.Fragment>
                        </tbody>
                    </Table>
                </CardBody>
            </Card>
        );
    }
}

TranslationFilesTable.propTypes = {
    listTranslationFiles: PropTypes.array,
    handleDownloadTranslationFile: PropTypes.func.isRequired,
    handleDeleteTranslationFile: PropTypes.func.isRequired,
};

TranslationFilesTable.defaultProps = {
    listTranslationFiles: []
};
