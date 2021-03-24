import React from 'react';
import {
    ButtonGroup,
    Button,
    Card,
    CardBody,
    Table
} from './../../components';
import PropTypes from 'prop-types';

export class LocalizationTable extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        const {listLocalizations} = this.props;
        return (
            <Card className="mb-3">
                <CardBody>
                    <Table className="mb-0" bordered responsive>
                        <thead>
                        <tr>
                            <th>ID</th>
                            <th>Language name</th>
                            <th>Localization code</th>
                            <th className="text-right">
                                Actions
                            </th>
                        </tr>
                        </thead>
                        <tbody>
                        <React.Fragment>
                            {
                                listLocalizations && listLocalizations.map((product, index) => (
                                        <tr key={index}>
                                            <td className="align-middle">
                                                {product.locale}
                                            </td>
                                            <td className="align-middle">
                                                {product.lang_name} {product.icon}
                                            </td>
                                            <td className="align-middle">
                                            <span className="text-inverse">
                                                {product.locale}
                                            </span>
                                            </td>
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

LocalizationTable.propTypes = {
    listLocalizations: PropTypes.array,
};
