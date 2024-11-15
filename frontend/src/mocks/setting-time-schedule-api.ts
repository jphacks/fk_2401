import { TimeScheduleResponse } from "@/types/api";

export function getTimeSchedules(houseID: number): TimeScheduleResponse[] {
  const nasuTimeSchedules: TimeScheduleResponse[] = [
    {
      start_time: "00:00",
      end_time: "02:30",
      workflows: [
        {
          id: 1,
          name: "高めの加温",
        },
        {
          id: 2,
          name: "加湿",
        },
      ],
    },
    {
      start_time: "03:00",
      end_time: "05:00",
      workflows: [
        {
          id: 1,
          name: "高めの加温",
        },
        {
          id: 3,
          name: "CO2供給",
        },
      ],
    },
    {
      start_time: "05:30",
      end_time: "07:30",
      workflows: [
        {
          id: 2,
          name: "加湿",
        },
        {
          id: 4,
          name: "換気",
        },
      ],
    },
    {
      start_time: "08:00",
      end_time: "10:00",
      workflows: [
        {
          id: 1,
          name: "高めの加温",
        },
        {
          id: 5,
          name: "照明",
        },
      ],
    },
    {
      start_time: "10:30",
      end_time: "12:30",
      workflows: [
        {
          id: 3,
          name: "CO2供給",
        },
        {
          id: 4,
          name: "換気",
        },
      ],
    },
    {
      start_time: "13:00",
      end_time: "15:00",
      workflows: [
        {
          id: 2,
          name: "加湿",
        },
        {
          id: 5,
          name: "照明",
        },
      ],
    },
  ];

  const hourensouTimeSchedule: TimeScheduleResponse[] = [
    {
      start_time: "06:00",
      end_time: "08:00",
      workflows: [
        {
          id: 1,
          name: "加温",
        },
        {
          id: 2,
          name: "加湿",
        },
      ],
    },
    {
      start_time: "08:30",
      end_time: "10:30",
      workflows: [
        {
          id: 3,
          name: "CO2供給",
        },
        {
          id: 4,
          name: "換気",
        },
      ],
    },
    {
      start_time: "11:00",
      end_time: "13:00",
      workflows: [
        {
          id: 1,
          name: "加温",
        },
        {
          id: 5,
          name: "照明",
        },
      ],
    },
    {
      start_time: "13:30",
      end_time: "15:30",
      workflows: [
        {
          id: 2,
          name: "加湿",
        },
        {
          id: 4,
          name: "換気",
        },
      ],
    },
    {
      start_time: "16:00",
      end_time: "18:00",
      workflows: [
        {
          id: 3,
          name: "CO2供給",
        },
        {
          id: 5,
          name: "照明",
        },
      ],
    },
  ];

  const timeScheduleMap: Map<number, TimeScheduleResponse[]> = new Map();

  timeScheduleMap.set(2, nasuTimeSchedules);
  timeScheduleMap.set(3, hourensouTimeSchedule);

  const result: TimeScheduleResponse[] = timeScheduleMap.get(houseID) || [];

  return result;
}
