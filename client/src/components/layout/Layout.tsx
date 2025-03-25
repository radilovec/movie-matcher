import style from "./style.module.css";
import { Header } from "../header/Header";
import { Footer } from "../footer/Footer";
import { Outlet } from "react-router-dom";

export const Layout: React.FC = () => {
  return (
    <>
      <Header />
      <main className={style.main}>
        <div className={style.wrapper}>
          <Outlet />
        </div>
      </main>
      <Footer />
    </>
  );
};
