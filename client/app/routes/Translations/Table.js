import React from 'react';
import BootstrapTable from 'react-bootstrap-table-next';
import ToolkitProvider from 'react-bootstrap-table2-toolkit';
import moment from 'moment';
import {
    Button,
    ButtonGroup,
    Row,
    Col
} from './../../components';
import * as PropTypes from 'prop-types';
import {Typeahead} from 'react-bootstrap-typeahead';

export class TranslationTable extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            chosenStatusID: 0
        };
    }

    openEditView = (id) => {
        const translation = this.props.listTranslations.find(o => o.id === id);
        this.props.history.push({
            pathname: `/translations/edit/${translation.id}`,
            state: translation,
        });
    };

    handleStatusTranslation = (e, id, status = 0) => {
        e.stopPropagation();
        this.props.updateStatusTranslationRequest(id, status);
    };

    createColumnDefinitions() {
        return [
            {
                dataField: 'localization',
                text: 'Localization',
                formatter: (cell) => {
                    if (cell !== undefined) {
                        return cell.locale;
                    } else {
                        return null;
                    }
                },
            }, {
                dataField: 'identifier',
                text: 'Identifier',
                formatter: (cell) => {
                    if (cell !== undefined) {
                        return cell.name;
                    } else {
                        return null;
                    }
                },
            }, {
                dataField: 'text',
                text: 'Text',
                formatter: (cell, row) => {
                    if (row.singular.length > 70) {
                        return row.singular.substring(0, 70) + '...';
                    } else {
                        return row.singular;
                    }
                },
            }, {
                dataField: 'status',
                text: 'Status',
                sort: true,
            }, {
                dataField: 'id',
                text: 'Action',
                sort: false,
                formatter: (cell, row) => {
                    const statusButtons = ['Hidden', 'Draft', 'Release'];
                    return (
                        <>
                            <ButtonGroup>
                                <Button onClick={this.openEditView.bind(this, row.id)}>
                                    Edit
                                </Button>
                            </ButtonGroup>
                            &nbsp;&nbsp;&nbsp;&nbsp;
                            <ButtonGroup>
                                {statusButtons.map((statusButton, i) => (
                                    <Button
                                        key={i}
                                        name={statusButton}
                                        onClick={
                                            (event) => this.handleStatusTranslation(event, row.id, i)
                                        }
                                        className={statusButton === row.status ? 'active' : ''}
                                        color="primary"
                                    >
                                        {statusButton}
                                    </Button>
                                ))}
                            </ButtonGroup>
                        </>
                    );
                },
            }
        ];
    }

    render() {
        const columnDefs = this.createColumnDefinitions();
        const {listTranslations, listLocalizations, chooseLocalization, handleChosenLocalization} = this.props;
        const expandRow = {
            renderer: row => (
                <Row>
                    <Col md={6}>
                        <dl className="row">
                            <dt className="col-sm-6 text-right">Description</dt>
                            <dd className="col-sm-6">{row.identifier.description}</dd>
                            <dt className="col-sm-6 text-right">Example text singular</dt>
                            <dd className="col-sm-6">
                                {row.identifier.text_singular.substring(0, 100) + '...'}
                            </dd>
                            <dt className="col-sm-6 text-right">Example text plural</dt>
                            <dd className="col-sm-6">
                                {row.identifier.text_plural.substring(0, 100) + '...'}
                            </dd>
                        </dl>
                    </Col>
                    <Col md={6}>
                        <dl className="row">
                            <dt className="col-sm-6 text-right">Created At</dt>
                            <dd className="col-sm-6">{moment.unix(row.created_at).format('DD-MMM-YYYY')}</dd>
                            <dt className="col-sm-6 text-right">Last Updated At</dt>
                            <dd className="col-sm-6">{moment.unix(row.created_at).format('DD-MMM-YYYY')}</dd>
                        </dl>
                    </Col>
                </Row>
            ),
            showExpandColumn: true,
            expandHeaderColumnRenderer: ({isAnyExpands}) => isAnyExpands ? (
                <i className="fa fa-angle-down fa-fw fa-lg text-muted"/>
            ) : (
                <i className="fa fa-angle-right fa-fw fa-lg text-muted"/>
            ),
            expandColumnRenderer: ({expanded}) => expanded ? (
                <i className="fa fa-angle-down fa-fw fa-lg text-muted"/>
            ) : (
                <i className="fa fa-angle-right fa-fw fa-lg text-muted"/>
            )
        };
        const defaultSelected = Object.keys(chooseLocalization).length === 0 ? [] : [chooseLocalization];
        return (
            <ToolkitProvider
                keyField="id"
                data={listTranslations}
                columns={columnDefs}
                search
            >
                {
                    props => (
                        <React.Fragment>
                            <div className="justify-content-end align-items-center mb-2">
                                <div className="d-flex my-1">
                                    <Typeahead
                                        clearButton
                                        id="list-localizations-for-id"
                                        defaultSelected={defaultSelected}
                                        labelKey="locale"
                                        options={listLocalizations}
                                        placeholder="Choose a localization..."
                                        onChange={handleChosenLocalization}
                                    />
                                </div>
                            </div>
                            <BootstrapTable
                                classes="table-responsive-lg"
                                bordered={true}
                                expandRow={expandRow}
                                responsive
                                hover
                                {...props.baseProps}
                            />
                        </React.Fragment>
                    )
                }
            </ToolkitProvider>
        );
    }
}

TranslationTable.propTypes = {
    listTranslations: PropTypes.array.isRequired,
    listLocalizations: PropTypes.array.isRequired,
    history: PropTypes.shape({
        push: PropTypes.func.isRequired,
    }).isRequired,
    handleChosenLocalization: PropTypes.func.isRequired,
    updateStatusTranslationRequest: PropTypes.func.isRequired,
    chooseLocalization: PropTypes.object,
};

TranslationTable.defaultProps = {
    listTranslations: [],
    listLocalizations: [],
    chooseLocalization: {},
};
