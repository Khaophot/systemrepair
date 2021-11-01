import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { RepairInformationsInterface } from "../models/IRepairInformation";

import { format } from 'date-fns'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function RepairInformations() {
  const classes = useStyles();
  const [repairinformations, setRepairInformations] = useState<RepairInformationsInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getRepairInformations = async () => {
    fetch(`${apiUrl}/repairinformations`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
            setRepairInformations(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getRepairInformations();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="lg">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลการแจ้งซ่อม
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="10%">
                  เลขห้อง
                </TableCell>
                <TableCell align="center" width="15%">
                  ชื่อลูกค้า
                </TableCell>
                <TableCell align="center" width="15%">
                  ชื่ออุปกรณ์
                </TableCell>
                <TableCell align="center" width="20%">
                  ปัญหาของอุปกรณ์
                </TableCell>
                <TableCell align="center" width="15%">
                  ความเร่งด่วน
                </TableCell>
                <TableCell align="center" width="40%">
                  วันที่และเวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {repairinformations.map((item: RepairInformationsInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.CheckIn.Room.Roomnumber}</TableCell>
                  <TableCell align="center">{item.CheckIn.Customer.Name}</TableCell>
                  <TableCell align="center">{item.Equipment.Name}</TableCell>
                  <TableCell align="center">{item.Problem.Value}</TableCell>
                  <TableCell align="center">{item.Urgency.Value}</TableCell>
                  <TableCell align="center">{format((new Date(item.Datetime)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default RepairInformations;