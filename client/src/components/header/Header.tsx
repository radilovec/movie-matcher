import { Button } from "../button/Button";
import style from "./style.module.css";

export const Header: React.FC = () => {
  return (
    <header className={style.header}>
      <div className={style.container}>
        <div className={style.wrapper}>
          <div className={style.logo}>
            <img src="img/logo.svg" alt="movie match" />
          </div>

          <div className={style.buttons}>
            <Button text="Sign In" bgColor="transparent" />
            <Button text="Login" />
          </div>
        </div>
      </div>
    </header>
  );
};
