import React from "react";
import style from "./style.module.css";

type Props = {
  text: string;
  bgColor?: string;
};

export const Button: React.FC<Props> = ({ text, bgColor }) => {
  return (
    <button className={style.btn} style={{ backgroundColor: bgColor }}>
      {text}
    </button>
  );
};
