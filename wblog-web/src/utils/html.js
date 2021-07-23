const htmlUtil = {
  /*
   * 文章概要用到了这个方法，去掉图片
   */
  trimImg(str) {
    str = str.replace(/<\s*img\s+.*?\/?\s*>/g, "");
    return str;
  },
  entity2HTML(str) {
    let arrEntities = {
      lt: "<",
      gt: ">",
      nbsp: " ",
      amp: "&",
      quot: '"',
    };
    let reg = /&(lt|gt|nbsp|amp|quot);/gi;
    return str.replace(reg, function (all, t) {
      return arrEntities[t];
    });
  },
};

export default htmlUtil;
