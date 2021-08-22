import {React, useState, useEffect} from 'react'
import { Col, Container, Form, Row, Spinner, Table } from 'react-bootstrap'
import { Arrow90degRight } from 'react-bootstrap-icons'

//main component used to handle state and call apis
const Restreamers = () => {
    //nned to pass restreamer data and status in each object before
    //passing to child components
    const list = [
        {
            name : "pippo",
            owner : "footters",
            status : "running"
        },
        {
            name : "pluto",
            owner : "lvf",
            status : "stopped"
        },
        {
            name : "paperino",
            owner : "demo",
            status : "running"
        }
    ]

    const [restreamers, setRestreamers] = useState(list)
    const [restreamersStatus, setRestreamersStatus] = useState([])

    //call every 10 secs for restreamer list and status
    // setInterval(() => {
    //     //chiamata axios lista e chiamata status
    //     //setRestreamers(value axios)
    //     //restreamersStatus(value axios)
    // }, 10000);

    return (
        <Container fluid>
            <Row>
                <Col md={4} className="border border-primary">
                    <RestreamersTable restreamers={restreamers}/>
                </Col>
                <Col md={8} className="border border-primary">
                col2
                </Col>
            </Row>
        </Container>
    )

    //render restreamer details card
    function renderCard() {
        
    }

}


//table component with search field
const RestreamersTable = (props) => {

    const [searchField, setSearchField] = useState("")
    const [restreamerSearchResults, setRestreamerSearchResults] = useState([])

     //searchfield filtering list
     useEffect(() => {
        const results = props.restreamers.filter(restr =>
          restr.name.includes(searchField) || restr.owner.includes(searchField)
        );

        setRestreamerSearchResults(results);
      }, [searchField, props.restreamers]); //rerender each restreamer list change and with searchfiled


    return(
    <div>
        <Form>
            <Form.Group className="mb-3" controlId="formName">
                <Form.Label>Search by Name or Owner</Form.Label>
                <Form.Control type="text" placeholder="Enter name or owner" value={searchField} onChange={(e) => setSearchField(e.target.value)} />
            </Form.Group>
        </Form>
        <Table class="table table-hover table-sm">  
            <thead>
                <tr>
                    <th>
                        #
                    </th>
                    <th>
                        Name
                    </th>
                    <th>
                        Owner
                    </th>
                    <th>
                        Status
                    </th>
                </tr>
            </thead>
            <tbody>
                {renderRows(restreamerSearchResults)}
            </tbody>
        </Table>
        </div>
    )

    //render multiple rows
    function renderRows(restreamers) {
        return restreamers.map(restr => {
            return <TableRow restreamer={restr} status={restr.status}/>
        }) 
    }

}


//single table row component
const TableRow = (props) => {
    return <tr className={props.status === "running" ? "table-success" : "table-active"}>
                    <td>
                        {props.status === "running" ?<Spinner size="sm" animation="border" variant="success" /> : <Arrow90degRight/>}
                    </td>
                    <td>
                        {props.restreamer.name}
                    </td>
                    <td>
                        {props.restreamer.owner}
                    </td>
                    <td>
                        {props.restreamer.status}
                    </td>
                </tr>
}

export default Restreamers
