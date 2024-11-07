-- name: CreateM304 :execlastid
INSERT INTO m304 (uecs_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetM304FromID :one
SELECT id, uecs_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name
FROM m304
WHERE id = ?;
