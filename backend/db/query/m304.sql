-- name: GetM304FromUecsDevice :many
SELECT id, uecs_id, mac_addr, dhcp_flg, ip_addr, net_mask, defgw, dns, vender_name, node_name, rly_0, rly_1, rly_2, rly_3, rly_4, rly_5, rly_6, rly_7
FROM m304
WHERE `rly_0` = ?
OR `rly_1` = ?
OR `rly_2` = ?
OR `rly_3` = ?
OR `rly_4` = ?
OR `rly_5` = ?
OR `rly_6` = ?
OR `rly_7` = ?;