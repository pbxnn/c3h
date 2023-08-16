insert into `c3h_data_info` (`name`, `key`, `desc`, `type`, `real_value`, `set_value`, `calc_value`, `low_limit`,
                             `high_limit`, `unit`)
values
-- ('DC402总开关','DC_402_Total', 'DC402总开关', 1, 0, 0, 0, 0, 0, ''),
-- ('DC402手动开关','DC_402_Manual', 'DC402手动开关', 1, 0, 0, 0, 0, 0, ''),
-- ('DC402AB选择','DC_402_AB', 'DC402AB选择', 1, 0, 0, 0, 0, 0, ''),
-- ('MAPD控制模式','MAPD_control_mode', '出口MAPD选择', 2, 0, 600.00, 0, 0, 0, ''),
-- ('干燥剂切换','drier', '干燥剂切换', 1, 0, 0, 0, 0, 0, 0, ''),
-- ('通讯中断报警','communication', '通讯中断报警', 1, 0, 0, 0, 0, 0, ''),
-- ('入口MAPD浓度','ingress_MAPD', '入口MAPD浓度', 2, 0, 0, 0, 0, 0, 'mol%'),
-- ('出口MAPD浓度','egress_MAPD', '出口MAPD浓度', 2, 0, 0, 0, 0, 0, 'mol%'),
-- ('入口C3流量','ingress_c3_flow', '入口C3流量', 2, 0, 0, 0, 0, 0, 'Kg/h'),
-- ('C3回流量','c3_back_flow', 'C3回流量', 2, 0, 0, 0, 0, 0, 'Kg/h'),
-- ('上层H2流量','upper_h2_flow', '上层H2流量', 2, 111.90, 112.00, 0, 50, 200, 'Kg/h'),
-- ('配氢比','h_ratio', '配氢比', 2, 0, 600.00, 0, 0, 0, '%'),
-- ('下层H2流量','lower_h2_flow', '下层H2流量', 2, 111.90, 112.00, 0, 50, 200, 'Kg/h'),
-- ('床测总H2流量','total_h2_flow', '床测总H2流量', 2, 111.90, 112.00, 0, 50, 200, 'Kg/h'),
-- ('AB入口温度','AB_ingress_temperature', 'AB入口温度', 2, 36.13, 38.18, 0, 0, 0, '°C'),
-- ('AB出口温度','AB_egress_temperature', 'AB出口温度', 2, 17.22, 0, 0, 0, 0, '°C'),
-- ('FA409压力','FA409_pressure', 'FA409压力', 2, 54.75, 0, 0, 0, 0, 'KPA'),
-- ('入口MAPD分析值','ingress_MAPD_analysis', '入口MAPD分析值', 2, 1.23, 0, 0, 0, 0, 'mol%'),
-- ('出口MAPD分析值','egress_MAPD_analysis', '出口MAPD分析值', 2, 400.00, 0, 0, 0, 0, 'mol%'),
-- ('入口丙烯分析值','ingress_c3h6_analysis', '入口丙烯分析值', 2, 89.20, 0, 0, 0, 0, 'mol%'),
-- ('出口丙烯分析值','egress_c3h6_analysis', '出口丙烯分析值', 2, 89.70, 0, 0, 0, 0, 'mol%'),
-- ('入口丙烷分析值','ingress_c3h8_analysis', '入口丙烷分析值', 2, 7.40, 0, 0, 0, 0, 'mol%'),
-- ('出口丙烷分析值','egress_c3h8_analysis', '出口丙烷分析值', 2, 8.02, 0, 0, 0, 0, 'mol%'),
-- ('氢炔比','h_alk', '氢炔比', 2, 1.05, 0, 0, 0, 0, 'mol/mol'),
-- ('选择性','selectivity', '选择性', 2, 42.72, 0, 0, 0, 0, '%'),
-- ('转化率','inversion_rate', '转化率', 2, 96.7456, 0, 0, 0, 0, '%');

('MAPD控制模式', 'R401S_MAPD_control_mode', 'MAPD控制模式', 2, 0, 600.00, 0, 0, 0, ''),
('APC主开关', 'R401S_APC_main_switch', 'APC主开关', 1, 0, 0, 0, 0, 0, ''),
('APC写入开关', 'R401S_APC_write_switch', 'APC写入开关', 1, 1, 1, 0, 0, 0, ''),
('通讯', 'R401S_communication', '通讯', 1, 1, 1, 0, 0, 0, ''),

('FIC-40012', 'R401S_cycle_materials_flow', '循环物料流量', 2, 20.50, 57.9834, 57.8838, 55.00, 80.00, 't/h'),
('FIC-40013', 'R401S_hydrogen_ratio', '配氢比', 2, 1.88300, 1.50639, 1.50906, 1.20000, 1.90000, ''),
('FIC-40025', 'R401S_ingress_temperature', '入口温度', 2, 42.50, 39.55, 39.658, 35.000, 43.000, '°C'),
('FIC-40013', 'R401S_hydrogen_flow', '氢气流量', 2, 118.78, 119.31, 00.00, 50.00, 200.00, 'kg/h'),
('FIC-40008', 'R401S_reactor_pressure', '反应器压力', 2, 30.20, 30.00, 31.00, 0.00, 0.00, 'MPA'),
('FIC-40012', 'R401S_cycle_materials_flow_control_switch', '循环物料流量控制开关', 1, 1, 1, 0, 0, 0, ''),
('FIC-40013', 'R401S_hydrogen_ratio_control_switch', '配氢比控制开关', 1, 1, 1, 0, 0, 0, ''),
('FIC-40025', 'R401S_ingress_temperature_control_switch', '入口温度控制开关', 1, 1, 1, 0, 0, 0, ''),

('FIC-40014A', 'R401S_mixed_total_materials_flow', '混合总物料流量', 2, 132.57, 132.57, 0, 0.00, 200.00, ''),
('FIC-40003', 'R401S_fresh_materials_flow', '新鲜物料流量', 2, 63.09, 63.09, 0, 0.00, 100.00, ''),
('FIC-40040A', 'R401S_egress_temperature', '出口温度', 2, 55.31, 55.31, 0, 0.00, 100.00, '');

('PBI-40007B', 'R401S_reactor_pressure_drop', '反应器压降', 2, 99.43, 0, 0, 0.00, 200.00, 'KPA'),
('TI-40037AA', 'R401S_up_bed_temperature_AA', '上层床层温度', 2, 41.33, 0, 0, 0.00, 100.00, '°C'),
('TI-40037AB', 'R401S_up_bed_temperature_AB', '上层床层温度', 2, 41.16, 0, 0, 0.00, 100.00, '°C'),
('TI-40038AA', 'R401S_mid_bed_temperature_AA', '中层床层温度', 2, 50.06, 0, 0, 0.00, 100.00, '°C'),
('TI-40038AB', 'R401S_mid_bed_temperature_AB', '中层床层温度', 2, 50.56, 0, 0, 0.00, 100.00, '°C'),
('TI-40039AA', 'R401S_down_bed_temperature_AA', '下层床层温度', 2, 54.59, 0, 0, 0.00, 100.00, '°C'),
('TI-40039AB', 'R401S_down_bed_temperature_AB', '下层床层温度', 2, 54.59, 0, 0, 0.00, 100.00, '°C'),


('催化剂性能参数', 'R401S_c3h6_selectivity', '丙烯选择性', 2, 89.54, 0, 0, 0, 0, '%'),
('催化剂性能参数', 'R401S_MAPD_inversion_rate', 'MAPD转化率', 2, 95.31, 0, 0, 0, 0, '%'),
('在线分析', 'R401S_ingress_MAPD_analysis', '入口MAPD', 2, 1.18, 0, 0, 0, 0, 'mol%'),
('在线分析', 'R401S_egress_MAPD_analysis', '出口MAPD', 2, 563.35, 0, 0, 0, 0, 'ppm'),
('在线分析', 'R401S_ingress_c3h6_analysis', '入口丙烯', 2, 87.55, 0, 0, 0, 0, 'mol%'),
('在线分析', 'R401S_egress_c3h6_analysis', '出口丙烯', 2, 90.27, 0, 0, 0, 0, 'mol%'),
('在线分析', 'R401S_ingress_c3h8_analysis', '入口丙烷', 2, 8.86, 0, 0, 0, 0, 'mol%'),
('在线分析', 'R401S_egress_c3h8_analysis', '出口丙烷', 2, 9.14, 0, 0, 0, 0, 'mol%'),
('在线分析', 'R401S_ingress_c4', '入口c4', 2, 0.00, 0, 0, 0, 0, 'mol%'),
('在线分析', 'R401S_egress_hydrogen', '出口氢气', 2, 0.32, 0, 0, 0, 0, 'mol%'),
