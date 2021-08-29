import React from 'react';
import {
    Button, Row, Col
} from './../../components';
import PropTypes from 'prop-types';
import ToolkitProvider from 'react-bootstrap-table2-toolkit';
import BootstrapTable from 'react-bootstrap-table-next';
import {ButtonGroup} from 'reactstrap';
import CustomSearch from '../../components/CustomSearch';

const sortCaret = (order) => {
    if (order) {
        return <i className={`fa fa-fw text-muted fa-sort-${order}`}/>;
    } else {
        return <i className="fa fa-fw fa-sort text-muted"/>;
    }
};

const defaultString = '';

export class IdentifierTable extends React.Component {
    createColumnDefinitions() {
        return [
            {
                dataField: 'name',
                text: 'Name',
                sort: true,
                sortCaret
            }, {
                dataField: 'category',
                text: 'Category',
                formatter: (cell) => {
                    return (cell !== undefined) ? cell.name : null;
                },
            }, {
                dataField: 'platforms',
                text: 'Platforms',
                formatter: (cell) => {
                    let icons = [];
                    cell.forEach(function (value) {
                        const v = value === 'web' ? 'chrome' : value;
                        icons.push(<i className={`fa fa-fw fa-${v}`} aria-hidden="true"/>);
                    });
                    return <div className="row">
                        {icons}
                    </div>;
                },
            }, {
                dataField: 'id',
                text: 'Action',
                sort: false,
                formatter: (cell) => {
                    return (
                        <ButtonGroup>
                            <Button onClick={() => this.props.handleChosenIdentifier(cell)} color="primary">
                                Update
                            </Button>
                            <Button onClick={() => this.props.deleteIdentifierRequest(cell)} color="danger">
                                Delete
                            </Button>
                        </ButtonGroup>
                    );
                },
            }
        ];
    }

    render() {
        const {listIdentifiers} = this.props;

        const columnDefs = this.createColumnDefinitions();

        const expandRow = {
            renderer: row => (
                <Row>
                    <Col md={6}>
                        <dl className="row">
                            <dt className="col-sm-6 text-right">Description</dt>
                            <dd className="col-sm-6">{row.description || defaultString}</dd>

                            <dt className="col-sm-6 text-right">Example Text</dt>
                            <dd className="col-sm-6">{row.example_text || defaultString}</dd>
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

        return (
            <ToolkitProvider
                keyField="id"
                data={listIdentifiers}
                columns={columnDefs}
                search
            >
                {
                    props => (
                        <React.Fragment>
                            <div className="d-flex justify-content-end align-items-center mb-2">
                                <h6 className="my-0">
                                    List of all identifiers
                                </h6>
                                <div className="d-flex ml-auto">
                                    <CustomSearch
                                        className="mr-2"
                                        {...props.searchProps}
                                    />
                                </div>
                            </div>
                            <BootstrapTable
                                classes="table-responsive-lg"
                                bordered={false}
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

IdentifierTable.propTypes = {
    listIdentifiers: PropTypes.array,
    handleChosenIdentifier: PropTypes.func.isRequired,
    deleteIdentifierRequest: PropTypes.func.isRequired,
};

IdentifierTable.defaultProps = {
    listIdentifiers: []
};
