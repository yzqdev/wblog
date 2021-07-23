import dayjs from 'dayjs'

const dateTool = {
  parse(dateStr) {
    let formatStr = "YYYY-MM-DDTHH:mm:ssZ";
    if (dateStr.indexOf(".") >= 0) {
      formatStr = "YYYY-MM-DDTHH:mm:ss.SSSSSSSSSZ";
    }
    return dayjs(dateStr, formatStr);
  },
  format(dayjs) {
    return dayjs.format("YYYY-MM-DDTHH:mm:ssZ");
  },
  formatYMD(dateStr) {
    let time = dateTool.parse(dateStr);
    return dayjs(time).format("YYYY-MM-DD");
  },
  formatYMDHM(dateStr) {
    let time = dateTool.parse(dateStr);
    return dayjs(time).format("YYYY-MM-DD HH:mm");
  },
  formatYMDHM2(dateStr) {
    let time = dateTool.parse(dateStr);
    console.log(time.year(), time.month(), time.date(), time.second());
    return dayjs(time).format("YYYY年MM月DD日 HH点mm分ss秒");
  },
  formatYMDHMS(dateStr) {
    let time = dateTool.parse(dateStr);
    return dayjs(time).format("YYYY-MM-DD HH:mm:ss");
  },
  getReplyTime(date) {
    let time = dateTool.parse(date).valueOf();
    let currentT = new Date().getTime();
    let diff = (currentT - time) / 1000;
    if (diff < 60) {
      return "刚刚";
    } else if (diff < 60 * 60) {
      return `${parseInt(diff / 60)}分钟前`;
    } else if (diff < 24 * 60 * 60) {
      return `${parseInt(diff / 60 / 60)}小时前`;
    } else if (diff < 7 * 24 * 60 * 60) {
      return `${parseInt(diff / 24 / 60 / 60)}天前`;
    } else {
      return dayjs(time).format("YYYY-MM-DD");
      // return `${parseInt(diff / 365 / 24 / 60 / 60)}年前`
    }
  },
};

export default dateTool;
