import {
  Box,
  Drawer,
  Link,
  List,
  ListItem,
  ListItemButton,
  Toolbar,
} from "@mui/material";

const drawerWidth = 240;

interface SidebarItem {
  title: string;
  path: string;
}

const sidebarItems: SidebarItem[] = [
  {
    title: "ハウス状態",
    path: "/",
  },
  {
    title: "デバイス設定",
    path: "/devices",
  },
  {
    title: "ワークフロー制御",
    path: "/workflow",
  },
];

const SideBar = () => {
  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        [`& .MuiDrawer-paper`]: {
          width: drawerWidth,
          boxSizing: "border-box",
          backgroundColor: "primary.light",
        },
      }}
    >
      <Toolbar />
      <Box sx={{ overflow: "auto" }}>
        <List>
          {sidebarItems.map((item, index) => (
            <ListItem key={index} disablePadding>
              <ListItemButton>
                {/* <ListItemIcon>
                  {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                </ListItemIcon> */}
                <Link
                  href={item.path}
                  underline="none"
                  color="inherit"
                  sx={{
                    color: "#000",
                    "&:hover": {
                      color: "#000",
                    },
                  }}
                >
                  {item.title}
                </Link>
              </ListItemButton>
            </ListItem>
          ))}
        </List>
      </Box>
    </Drawer>
  );
};

export default SideBar;
