import { ThemeProvider } from "@mui/material/styles";
import Header from "../components/header";
import Sidebar from "../components/sidebar";
import { uecsTheme } from "../styles/theme";

export function Navigation() {
  return (
    <>
      <ThemeProvider theme={uecsTheme}>
        <Header></Header>
        <Sidebar></Sidebar>
      </ThemeProvider>
    </>
  );
}
