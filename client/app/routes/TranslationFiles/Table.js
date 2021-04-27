import React from 'react';
import {
    ButtonGroup,
    Button, Card, CardBody, Table
} from './../../components';
import PropTypes from "prop-types";


export class GenerationTable extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        const {listLocalization} = this.props;
        return (
            <Card className="mb-3">
                <CardBody>
                    <Table className="mb-0" bordered responsive>
                        <thead>
                        <tr>
                            <th>Language name</th>
                            <th>Localization name</th>
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
                                listLocalization.map((product) => (
                                        <tr key={product.id}>
                                            <td className="align-middle">
                                            <span className="text-inverse">
                                                {product.local_name}
                                            </span>
                                            </td>
                                            <td className="align-middle">
                                                {product.lang_name}
                                            </td>
                                            <td className="align-middle">
                                                {product.lcid}
                                            </td>
                                            <td className="align-middle">
                                                {product.lcid}
                                            </td>
                                            <td className="text-right">
                                                <ButtonGroup>
                                                    <Button
                                                        color="link"
                                                        className="text-decoration-none"
                                                    >
                                                        <i className="fa fa-download"></i>
                                                    </Button>
                                                    <Button
                                                        color="link"
                                                        className="text-decoration-none"
                                                    >
                                                        <i className="fa fa-close"></i>
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

GenerationTable.propTypes = {
    listLocalization: PropTypes.array.isRequired,
};
