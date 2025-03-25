import React from "react";
import style from "./style.module.css";

type Props = {
  param: string;
  val: string;
};

export const InfBox: React.FC<Props> = ({ param, val }) => {
  return (
    <div className={style.box}>
      <span className={style.param}>{param}</span>
      <span className={style.value}>{val}</span>
    </div>
  );
};
