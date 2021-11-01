import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

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

function Home() {
  const classes = useStyles();

  return (
    <div>

      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบแจ้งซ่อม</h1>
        <h4>Requirements</h4>
        <p>
              ภายในระบบโรงแรมของโรงแรมอัครเดช เป็นระบบที่ให้ผู้ใช้ระบบซึ่งเป็นลูกค้าสามารถเลือกอุปกรณ์
          ภายในห้องพักของตนเองภายในโรงแรมเพื่อแจ้งต่อกับระบบว่าต้องการแจ้งซ่อมอุปกรณ์ภายในห้องพัก 
          โดยการแจ้งซ่อมจะประกอบไปด้วย หมายเลขการเช็คอิน (Check-in) ของตัวเองและเมื่อเลือกแล้ว
          ระบบจะทำการเลือกห้องพักของตนเองอัตโนมัติและมีชื่ออุปกรณ์ที่ต้องการจะแจ้งซ่อม ความเร่งด่วนที่ต้องการใช้งาน
          วันที่และเวลา และปัญหาของอุปกรณ์ที่ต้องการจะซ่อม และเมื่อแจ้งซ่อมกับระบบแล้วจะทำการบันทึกข้อมูลการ
          แจ้งซ่อมเอาไว้
        </p>
      </Container>
    </div>
  );
}
export default Home;