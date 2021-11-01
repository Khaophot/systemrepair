import React from "react";
import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import MenuItem from "@material-ui/core/MenuItem";

import { CheckInsInterface } from "../models/ICheckIn";
import { RoomsInterface } from "../models/IRoom";
import { EquipmentsInterface } from "../models/IEquipment";
import { ProblemsInterface } from "../models/IProblem";
import { UrgenciesInterface } from "../models/IUrgency";

import {
  MuiPickersUtilsProvider,
  KeyboardDatePicker,
} from "@material-ui/pickers";

import DateFnsUtils from "@date-io/date-fns";
import { RepairInformationsInterface } from "../models/IRepairInformation";


function Alert(props: AlertProps) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: { flexGrow: 1 },
    container: { marginTop: theme.spacing(2) },
    paper: { padding: theme.spacing(2), color: theme.palette.text.secondary },
  })
);

function RepairInformationCreate() {

  const classes = useStyles();
  const [selectedDate, setSelectedDate] = React.useState<Date | null>(new Date());
  const [checkins, setCheckIns] = React.useState<CheckInsInterface[]>([]);
  const [room, setRoom] = React.useState<RoomsInterface>();
  const [equipments, setEquipments] = React.useState<EquipmentsInterface[]>([]);
  const [problems, setProblems] = React.useState<ProblemsInterface[]>([]);
  const [urgencies, setUrgencies] = React.useState<UrgenciesInterface[]>([]);
  const [repairInformation, setRepairInformation] = React.useState<Partial<RepairInformationsInterface>>({});
  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);
  const apiUrl = "http://localhost:8080";

  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json"
    },
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>
  ) => {
    const name = event.target.name as keyof typeof RepairInformationCreate;
    setRepairInformation({
      ...repairInformation,
      [name]: event.target.value,
    });

    // เมื่อเลือก checkin แล้ว ห้องจะเด้ง
    
    if (event.target.name === "CheckInID") {
      let checkin = checkins.find(checkIn => checkIn.ID === event.target.value);
      setRoom(checkin?.Room);
    }
  };

  // เรียกใช้คำสั่ง GET จาก backend โดยผ่าน requestOptions

  const getCheckIns = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/check_in/reserved/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCheckIns(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getEquipment = async () => {
    fetch(`${apiUrl}/equipments`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setEquipments(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getProblem = async () => {
    fetch(`${apiUrl}/problems`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setProblems(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getUrgency = async () => {
    fetch(`${apiUrl}/urgencies`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setUrgencies(res.data);
        } else {
          console.log("else");
        }
      });
  };

  // ใช้คำสั่ง useEffect เพื่อดึงข้อมูลจาก API และรีเฟรชข้อมูลนั้นเป็นระยะ

  useEffect(() => {
    getCheckIns();
    getEquipment();
    getProblem();
    getUrgency();
  }, []);

  // คำสั่ง convertType เปลี่ยน Type ของ output 
  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      CheckInID: convertType(repairInformation.CheckInID),
      EquipmentID: convertType(repairInformation.EquipmentID),
      ProblemID: convertType(repairInformation.ProblemID),
      UrgencyID: convertType(repairInformation.UrgencyID),
      DateTime: selectedDate,

    };

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/repairinformations`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
        }
      });
  }

  return (
    
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ระบบแจ้งซ่อม
            </Typography>
          </Box>
        </Box>
        <Divider />

        <Grid container spacing={6} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>หมายเลขเช็คอิน</p>
              <Select
                value={repairInformation.CheckInID}
                onChange={handleChange}
                inputProps={{
                  name: "CheckInID",
                }}
              >
                <MenuItem aria-label="None" value="">
                  -กรุณาเลือกหมายเลขเช็คอิน-
                </MenuItem>
                {checkins.map((item: CheckInsInterface) => (
                  <MenuItem value={item.ID} key={item.ID}>
                    {item.ID}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>

            <FormControl fullWidth variant="outlined">
              <p>เลขห้องพัก</p>

              {<Select
                native
                disabled
                value={room?.Roomnumber}
              >
                <option aria-label="None" value="">
                  {room?.Roomnumber}
                </option>
              </Select>}

            </FormControl>

          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ชื่ออุปกรณ์</p>

              <Select
                value={repairInformation.EquipmentID}
                onChange={handleChange}
                inputProps={{
                  name: "EquipmentID",
                }}
              >
                <MenuItem aria-label="None" value="">
                  -กรุณาเลือกชื่อของอุปกรณ์-
                </MenuItem>
                {equipments.map((item: EquipmentsInterface) => (
                  <MenuItem value={item.ID} key={item.ID}>
                    {item.Name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ปัญหาของอุปกรณ์</p>
              <Select
                value={repairInformation.ProblemID}
                onChange={handleChange}
                inputProps={{
                  name: "ProblemID",
                }}
              >
                <MenuItem aria-label="None" value="">
                  -กรุณาเลือกปัญหาของอุปกรณ์-
                </MenuItem>
                {problems.map((item: ProblemsInterface) => (
                  <MenuItem value={item.ID} key={item.ID}>
                    {item.Value}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ความเร่งด่วน</p>
              <Select
                native
                value={repairInformation.UrgencyID}
                onChange={handleChange}
                inputProps={{
                  name: "UrgencyID",
                }}
              >
                <option aria-label="None" value="">
                  -กรุณาเลือกความเร่งด่วน-
                </option>
                {urgencies.map((item: UrgenciesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Value}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
          </Grid>

          <Grid item xs={3}>
            <p>วันที่และเวลา</p>
          </Grid>

          <Grid item xs={3}>
            <MuiPickersUtilsProvider utils={DateFnsUtils}>
              <KeyboardDatePicker
                name="DateTime"
                format="yyyy-MM-dd hh:mm"
                value={selectedDate}
                onChange={handleDateChange}
                KeyboardButtonProps={{
                  "aria-label": "change date",
                }}
              />
            </MuiPickersUtilsProvider>
          </Grid>

          <Grid item xs={12}>

            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              ยืนยัน
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default RepairInformationCreate;