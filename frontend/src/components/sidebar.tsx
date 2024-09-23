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

const SideBar = () => {
  const ListItems: string[] = [
    "項目1",
    "項目2",
    "項目3",
    "項目4",
  ]

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
          {ListItems.map((item) => (
            <ListItem key={item} disablePadding>
              <ListItemButton>
                {/* <ListItemIcon>
                  {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                </ListItemIcon> */}
                <Link
                  href="/" 
                  underline="none" 
                  color="inherit"
                  sx={{
                    color: "#000",
                    "&:hover": {
                      color: "#000", 
                    },
                  }}
                >
                  {item}
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
