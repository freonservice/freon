import React from 'react';
import {
    Container,
    Row,
    Card,
    CardBody,
    Col, CardTitle, CardDeck, Table
} from '../../components';

import {HeaderMain} from '../components/HeaderMain';
import {connect} from 'react-redux';
import {statRequest} from '../../redux/stat/actions';
import * as PropTypes from 'prop-types';

export class Dashboard extends React.Component {
    constructor(props) {
        super(props);

        this.props.statRequest();
    }

    render() {
        const {
            count_users,
            count_categories,
            count_identifiers,
            count_localizations,
            stat_completed_translations
        } = this.props.stat;
        console.log(typeof stat_completed_translations);
        return (
            <Container>
                <Row className="mb-5">
                    <Col lg={12}>
                        <HeaderMain
                            title="Dashboard"
                            className="mb-4 mb-lg-4"
                        />
                    </Col>
                    <Col lg={3}>
                        <Card className="mb-3">
                            <CardBody>
                                <CardTitle tag="h6" className="mb-4">
                                    Active Users
                                </CardTitle>
                                <div>
                                    <div className="mb-3">
                                        <h2>{count_users}</h2>
                                    </div>
                                </div>
                            </CardBody>
                        </Card>
                    </Col>
                    <Col lg={3}>
                        <Card className="mb-3">
                            <CardBody>
                                <CardTitle tag="h6" className="mb-4">
                                    Count Localizations
                                </CardTitle>
                                <div>
                                    <div className="mb-3">
                                        <h2>{count_localizations}</h2>
                                    </div>
                                </div>
                            </CardBody>
                        </Card>
                    </Col>
                    <Col lg={3}>
                        <Card className="mb-3">
                            <CardBody>
                                <CardTitle tag="h6" className="mb-4">
                                    Count Identifiers
                                </CardTitle>
                                <div>
                                    <div className="mb-3">
                                        <h2>{count_identifiers}</h2>
                                    </div>
                                </div>
                            </CardBody>
                        </Card>
                    </Col>
                    <Col lg={3}>
                        <Card className="mb-3">
                            <CardBody>
                                <CardTitle tag="h6" className="mb-4">
                                    Count Categories
                                </CardTitle>
                                <div>
                                    <div className="mb-3">
                                        <h2>{count_categories}</h2>
                                    </div>
                                </div>
                            </CardBody>
                        </Card>
                    </Col>
                    <Col lg={6}>
                        <CardDeck>
                            <Card className="mb-3">
                                <CardBody>
                                    <CardTitle className="mb-1 d-flex">
                                        <h6>Stat Translations</h6>
                                    </CardTitle>
                                </CardBody>
                                <Table responsive striped className="mb-0">
                                    <thead>
                                    <tr>
                                        <th className="bt-0">Translation</th>
                                        <th className="bt-0">Completed</th>
                                        <th className="bt-0 text-right">Action</th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    {
                                        stat_completed_translations && stat_completed_translations.map(t => (
                                            <tr key={t.lang_name}>
                                                <td className="align-middle">
                                                    <span className="text-inverse">{t.lang_name}</span>
                                                </td>
                                                <td className="align-middle">
                                                    {t.percentage}%
                                                </td>
                                                <td className="align-middle text-right text-nowrap">
                                                    <a href="#" className="text-decoration-none">View <i
                                                        className="fa fa-angle-right"/></a>
                                                </td>
                                            </tr>
                                        ))
                                    }
                                    </tbody>
                                </Table>
                            </Card>
                        </CardDeck>
                    </Col>
                </Row>
            </Container>);
    }
}

Dashboard.propTypes = {
    stat: PropTypes.object,
    errorMsg: PropTypes.string,
    statRequest: PropTypes.func.isRequired,
};

Dashboard.defaultTypes = {
    stat: {
        stat_completed_translations: []
    }
};

const mapStateToProps = (state) => ({
    stat: state.stat.stat,
    errorMsg: state.stat.error,
});

const mapDispatchToProps = {
    statRequest
};

export default connect(mapStateToProps, mapDispatchToProps)(Dashboard);