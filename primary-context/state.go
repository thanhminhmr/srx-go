package primary_context

type HistoryState struct {
	nextIfFirst  uint8
	nextIfMiss   uint8
	nextIfSecond uint8
	nextIfThird  uint8
	firstCount   uint8
}

func NewHistoryState(
	firstCount uint8,
	nextIfFirst uint8,
	nextIfSecond uint8,
	nextIfThird uint8,
	nextIfMiss uint8,
) HistoryState {
	return HistoryState{
		nextIfFirst:  nextIfFirst,
		nextIfMiss:   nextIfMiss,
		nextIfSecond: nextIfSecond,
		nextIfThird:  nextIfThird,
		firstCount:   firstCount,
	}
}

var stateTable = [...]HistoryState{
	NewHistoryState(0, 27, 0, 0, 0),      //   0,  0,  0,  0
	NewHistoryState(0, 28, 1, 27, 0),     //   1,  0,  0,  1
	NewHistoryState(0, 28, 1, 54, 0),     //   2,  0,  0,  2
	NewHistoryState(0, 29, 2, 74, 0),     //   3,  0,  0,  3
	NewHistoryState(0, 31, 27, 1, 1),     //   4,  0,  1,  0
	NewHistoryState(0, 32, 28, 28, 1),    //   5,  0,  1,  1
	NewHistoryState(0, 32, 28, 55, 1),    //   6,  0,  1,  2
	NewHistoryState(0, 33, 29, 75, 1),    //   7,  0,  1,  3
	NewHistoryState(0, 31, 54, 1, 1),     //   8,  0,  2,  0
	NewHistoryState(0, 32, 55, 28, 1),    //   9,  0,  2,  1
	NewHistoryState(0, 32, 55, 55, 1),    //  10,  0,  2,  2
	NewHistoryState(0, 33, 56, 75, 1),    //  11,  0,  2,  3
	NewHistoryState(0, 35, 74, 2, 2),     //  12,  0,  3,  0
	NewHistoryState(0, 36, 75, 29, 2),    //  13,  0,  3,  1
	NewHistoryState(0, 36, 75, 56, 2),    //  14,  0,  3,  2
	NewHistoryState(0, 37, 76, 76, 2),    //  15,  0,  3,  3
	NewHistoryState(0, 39, 90, 3, 3),     //  16,  0,  4,  0
	NewHistoryState(0, 40, 91, 30, 3),    //  17,  0,  4,  1
	NewHistoryState(0, 40, 91, 57, 3),    //  18,  0,  4,  2
	NewHistoryState(0, 41, 92, 77, 3),    //  19,  0,  4,  3
	NewHistoryState(0, 43, 102, 3, 3),    //  20,  0,  5,  0
	NewHistoryState(0, 44, 103, 30, 3),   //  21,  0,  5,  1
	NewHistoryState(0, 44, 103, 57, 3),   //  22,  0,  5,  2
	NewHistoryState(0, 47, 110, 3, 3),    //  23,  0,  6,  0
	NewHistoryState(0, 48, 111, 30, 3),   //  24,  0,  6,  1
	NewHistoryState(0, 50, 116, 3, 3),    //  25,  0,  7,  0
	NewHistoryState(0, 51, 117, 30, 3),   //  26,  0,  7,  1
	NewHistoryState(1, 54, 4, 4, 4),      //  27,  1,  0,  0
	NewHistoryState(1, 55, 5, 31, 4),     //  28,  1,  0,  1
	NewHistoryState(1, 55, 5, 58, 4),     //  29,  1,  0,  2
	NewHistoryState(1, 56, 6, 78, 4),     //  30,  1,  0,  3
	NewHistoryState(1, 58, 31, 5, 5),     //  31,  1,  1,  0
	NewHistoryState(1, 59, 32, 32, 5),    //  32,  1,  1,  1
	NewHistoryState(1, 59, 32, 59, 5),    //  33,  1,  1,  2
	NewHistoryState(1, 60, 33, 79, 5),    //  34,  1,  1,  3
	NewHistoryState(1, 58, 58, 5, 5),     //  35,  1,  2,  0
	NewHistoryState(1, 59, 59, 32, 5),    //  36,  1,  2,  1
	NewHistoryState(1, 59, 59, 59, 5),    //  37,  1,  2,  2
	NewHistoryState(1, 60, 60, 79, 5),    //  38,  1,  2,  3
	NewHistoryState(1, 62, 78, 6, 6),     //  39,  1,  3,  0
	NewHistoryState(1, 63, 79, 33, 6),    //  40,  1,  3,  1
	NewHistoryState(1, 63, 79, 60, 6),    //  41,  1,  3,  2
	NewHistoryState(1, 64, 80, 80, 6),    //  42,  1,  3,  3
	NewHistoryState(1, 65, 93, 7, 7),     //  43,  1,  4,  0
	NewHistoryState(1, 66, 94, 34, 7),    //  44,  1,  4,  1
	NewHistoryState(1, 66, 94, 61, 7),    //  45,  1,  4,  2
	NewHistoryState(1, 67, 95, 81, 7),    //  46,  1,  4,  3
	NewHistoryState(1, 68, 104, 7, 7),    //  47,  1,  5,  0
	NewHistoryState(1, 69, 105, 34, 7),   //  48,  1,  5,  1
	NewHistoryState(1, 69, 105, 61, 7),   //  49,  1,  5,  2
	NewHistoryState(1, 70, 112, 7, 7),    //  50,  1,  6,  0
	NewHistoryState(1, 71, 113, 34, 7),   //  51,  1,  6,  1
	NewHistoryState(1, 72, 118, 7, 7),    //  52,  1,  7,  0
	NewHistoryState(1, 73, 119, 34, 7),   //  53,  1,  7,  1
	NewHistoryState(2, 74, 8, 8, 8),      //  54,  2,  0,  0
	NewHistoryState(2, 75, 9, 35, 8),     //  55,  2,  0,  1
	NewHistoryState(2, 75, 9, 62, 8),     //  56,  2,  0,  2
	NewHistoryState(2, 76, 10, 82, 8),    //  57,  2,  0,  3
	NewHistoryState(2, 78, 35, 9, 9),     //  58,  2,  1,  0
	NewHistoryState(2, 79, 36, 36, 9),    //  59,  2,  1,  1
	NewHistoryState(2, 79, 36, 63, 9),    //  60,  2,  1,  2
	NewHistoryState(2, 80, 37, 83, 9),    //  61,  2,  1,  3
	NewHistoryState(2, 78, 62, 9, 9),     //  62,  2,  2,  0
	NewHistoryState(2, 79, 63, 36, 9),    //  63,  2,  2,  1
	NewHistoryState(2, 79, 63, 63, 9),    //  64,  2,  2,  2
	NewHistoryState(2, 82, 82, 10, 10),   //  65,  2,  3,  0
	NewHistoryState(2, 83, 83, 37, 10),   //  66,  2,  3,  1
	NewHistoryState(2, 83, 83, 64, 10),   //  67,  2,  3,  2
	NewHistoryState(2, 84, 96, 11, 11),   //  68,  2,  4,  0
	NewHistoryState(2, 85, 97, 38, 11),   //  69,  2,  4,  1
	NewHistoryState(2, 86, 106, 11, 11),  //  70,  2,  5,  0
	NewHistoryState(2, 87, 107, 38, 11),  //  71,  2,  5,  1
	NewHistoryState(2, 88, 114, 11, 11),  //  72,  2,  6,  0
	NewHistoryState(2, 89, 115, 38, 11),  //  73,  2,  6,  1
	NewHistoryState(3, 90, 12, 12, 12),   //  74,  3,  0,  0
	NewHistoryState(3, 91, 13, 39, 12),   //  75,  3,  0,  1
	NewHistoryState(3, 91, 13, 65, 12),   //  76,  3,  0,  2
	NewHistoryState(3, 92, 14, 84, 12),   //  77,  3,  0,  3
	NewHistoryState(3, 93, 39, 13, 13),   //  78,  3,  1,  0
	NewHistoryState(3, 94, 40, 40, 13),   //  79,  3,  1,  1
	NewHistoryState(3, 94, 40, 66, 13),   //  80,  3,  1,  2
	NewHistoryState(3, 95, 41, 85, 13),   //  81,  3,  1,  3
	NewHistoryState(3, 93, 65, 13, 13),   //  82,  3,  2,  0
	NewHistoryState(3, 94, 66, 40, 13),   //  83,  3,  2,  1
	NewHistoryState(3, 96, 84, 14, 14),   //  84,  3,  3,  0
	NewHistoryState(3, 97, 85, 41, 14),   //  85,  3,  3,  1
	NewHistoryState(3, 98, 98, 15, 15),   //  86,  3,  4,  0
	NewHistoryState(3, 99, 99, 42, 15),   //  87,  3,  4,  1
	NewHistoryState(3, 100, 108, 15, 15), //  88,  3,  5,  0
	NewHistoryState(3, 101, 109, 42, 15), //  89,  3,  5,  1
	NewHistoryState(4, 102, 16, 16, 16),  //  90,  4,  0,  0
	NewHistoryState(4, 103, 17, 43, 16),  //  91,  4,  0,  1
	NewHistoryState(4, 103, 17, 68, 16),  //  92,  4,  0,  2
	NewHistoryState(4, 104, 43, 17, 17),  //  93,  4,  1,  0
	NewHistoryState(4, 105, 44, 44, 17),  //  94,  4,  1,  1
	NewHistoryState(4, 105, 44, 69, 17),  //  95,  4,  1,  2
	NewHistoryState(4, 104, 68, 17, 17),  //  96,  4,  2,  0
	NewHistoryState(4, 105, 69, 44, 17),  //  97,  4,  2,  1
	NewHistoryState(4, 106, 86, 18, 18),  //  98,  4,  3,  0
	NewHistoryState(4, 107, 87, 45, 18),  //  99,  4,  3,  1
	NewHistoryState(4, 108, 100, 19, 19), // 100,  4,  4,  0
	NewHistoryState(4, 109, 101, 46, 19), // 101,  4,  4,  1
	NewHistoryState(5, 110, 20, 20, 20),  // 102,  5,  0,  0
	NewHistoryState(5, 111, 21, 47, 20),  // 103,  5,  0,  1
	NewHistoryState(5, 112, 47, 21, 21),  // 104,  5,  1,  0
	NewHistoryState(5, 113, 48, 48, 21),  // 105,  5,  1,  1
	NewHistoryState(5, 112, 70, 21, 21),  // 106,  5,  2,  0
	NewHistoryState(5, 113, 71, 48, 21),  // 107,  5,  2,  1
	NewHistoryState(5, 114, 88, 22, 22),  // 108,  5,  3,  0
	NewHistoryState(5, 115, 89, 49, 22),  // 109,  5,  3,  1
	NewHistoryState(6, 116, 23, 23, 23),  // 110,  6,  0,  0
	NewHistoryState(6, 117, 24, 50, 23),  // 111,  6,  0,  1
	NewHistoryState(6, 118, 50, 24, 24),  // 112,  6,  1,  0
	NewHistoryState(6, 119, 51, 51, 24),  // 113,  6,  1,  1
	NewHistoryState(6, 118, 72, 24, 24),  // 114,  6,  2,  0
	NewHistoryState(6, 119, 73, 51, 24),  // 115,  6,  2,  1
	NewHistoryState(7, 120, 25, 25, 25),  // 116,  7,  0,  0
	NewHistoryState(7, 121, 26, 52, 25),  // 117,  7,  0,  1
	NewHistoryState(7, 122, 52, 26, 26),  // 118,  7,  1,  0
	NewHistoryState(7, 123, 53, 53, 26),  // 119,  7,  1,  1
	NewHistoryState(8, 124, 25, 25, 25),  // 120,  8,  0,  0
	NewHistoryState(8, 125, 26, 52, 25),  // 121,  8,  0,  1
	NewHistoryState(8, 126, 52, 26, 26),  // 122,  8,  1,  0
	NewHistoryState(8, 127, 53, 53, 26),  // 123,  8,  1,  1
	NewHistoryState(9, 128, 25, 25, 25),  // 124,  9,  0,  0
	NewHistoryState(9, 129, 26, 52, 25),  // 125,  9,  0,  1
	NewHistoryState(9, 130, 52, 26, 26),  // 126,  9,  1,  0
	NewHistoryState(9, 131, 53, 53, 26),  // 127,  9,  1,  1
	NewHistoryState(10, 132, 25, 25, 25), // 128, 10,  0,  0
	NewHistoryState(10, 133, 26, 52, 25), // 129, 10,  0,  1
	NewHistoryState(10, 134, 52, 26, 26), // 130, 10,  1,  0
	NewHistoryState(10, 135, 53, 53, 26), // 131, 10,  1,  1
	NewHistoryState(11, 136, 25, 25, 25), // 132, 11,  0,  0
	NewHistoryState(11, 137, 26, 52, 25), // 133, 11,  0,  1
	NewHistoryState(11, 138, 52, 26, 26), // 134, 11,  1,  0
	NewHistoryState(11, 139, 53, 53, 26), // 135, 11,  1,  1
	NewHistoryState(12, 140, 25, 25, 25), // 136, 12,  0,  0
	NewHistoryState(12, 141, 26, 52, 25), // 137, 12,  0,  1
	NewHistoryState(12, 142, 52, 26, 26), // 138, 12,  1,  0
	NewHistoryState(12, 143, 53, 53, 26), // 139, 12,  1,  1
	NewHistoryState(13, 144, 25, 25, 25), // 140, 13,  0,  0
	NewHistoryState(13, 145, 26, 52, 25), // 141, 13,  0,  1
	NewHistoryState(13, 146, 52, 26, 26), // 142, 13,  1,  0
	NewHistoryState(13, 147, 53, 53, 26), // 143, 13,  1,  1
	NewHistoryState(14, 148, 25, 25, 25), // 144, 14,  0,  0
	NewHistoryState(14, 149, 26, 52, 25), // 145, 14,  0,  1
	NewHistoryState(14, 150, 52, 26, 26), // 146, 14,  1,  0
	NewHistoryState(14, 151, 53, 53, 26), // 147, 14,  1,  1
	NewHistoryState(15, 152, 25, 25, 25), // 148, 15,  0,  0
	NewHistoryState(15, 153, 26, 52, 25), // 149, 15,  0,  1
	NewHistoryState(15, 154, 52, 26, 26), // 150, 15,  1,  0
	NewHistoryState(15, 155, 53, 53, 26), // 151, 15,  1,  1
	NewHistoryState(16, 156, 25, 25, 25), // 152, 16,  0,  0
	NewHistoryState(16, 157, 26, 52, 25), // 153, 16,  0,  1
	NewHistoryState(16, 158, 52, 26, 26), // 154, 16,  1,  0
	NewHistoryState(16, 159, 53, 53, 26), // 155, 16,  1,  1
	NewHistoryState(17, 160, 25, 25, 25), // 156, 17,  0,  0
	NewHistoryState(17, 161, 26, 52, 25), // 157, 17,  0,  1
	NewHistoryState(17, 162, 52, 26, 26), // 158, 17,  1,  0
	NewHistoryState(17, 163, 53, 53, 26), // 159, 17,  1,  1
	NewHistoryState(18, 164, 25, 25, 25), // 160, 18,  0,  0
	NewHistoryState(18, 165, 26, 52, 25), // 161, 18,  0,  1
	NewHistoryState(18, 166, 52, 26, 26), // 162, 18,  1,  0
	NewHistoryState(18, 167, 53, 53, 26), // 163, 18,  1,  1
	NewHistoryState(19, 168, 25, 25, 25), // 164, 19,  0,  0
	NewHistoryState(19, 169, 26, 52, 25), // 165, 19,  0,  1
	NewHistoryState(19, 170, 52, 26, 26), // 166, 19,  1,  0
	NewHistoryState(19, 171, 53, 53, 26), // 167, 19,  1,  1
	NewHistoryState(20, 172, 25, 25, 25), // 168, 20,  0,  0
	NewHistoryState(20, 173, 26, 52, 25), // 169, 20,  0,  1
	NewHistoryState(20, 174, 52, 26, 26), // 170, 20,  1,  0
	NewHistoryState(20, 175, 53, 53, 26), // 171, 20,  1,  1
	NewHistoryState(21, 176, 25, 25, 25), // 172, 21,  0,  0
	NewHistoryState(21, 177, 26, 52, 25), // 173, 21,  0,  1
	NewHistoryState(21, 178, 52, 26, 26), // 174, 21,  1,  0
	NewHistoryState(21, 179, 53, 53, 26), // 175, 21,  1,  1
	NewHistoryState(22, 180, 25, 25, 25), // 176, 22,  0,  0
	NewHistoryState(22, 181, 26, 52, 25), // 177, 22,  0,  1
	NewHistoryState(22, 182, 52, 26, 26), // 178, 22,  1,  0
	NewHistoryState(22, 183, 53, 53, 26), // 179, 22,  1,  1
	NewHistoryState(23, 184, 25, 25, 25), // 180, 23,  0,  0
	NewHistoryState(23, 185, 26, 52, 25), // 181, 23,  0,  1
	NewHistoryState(23, 186, 52, 26, 26), // 182, 23,  1,  0
	NewHistoryState(23, 187, 53, 53, 26), // 183, 23,  1,  1
	NewHistoryState(24, 188, 25, 25, 25), // 184, 24,  0,  0
	NewHistoryState(24, 189, 26, 52, 25), // 185, 24,  0,  1
	NewHistoryState(24, 190, 52, 26, 26), // 186, 24,  1,  0
	NewHistoryState(24, 191, 53, 53, 26), // 187, 24,  1,  1
	NewHistoryState(25, 192, 25, 25, 25), // 188, 25,  0,  0
	NewHistoryState(25, 193, 26, 52, 25), // 189, 25,  0,  1
	NewHistoryState(25, 194, 52, 26, 26), // 190, 25,  1,  0
	NewHistoryState(25, 195, 53, 53, 26), // 191, 25,  1,  1
	NewHistoryState(26, 196, 25, 25, 25), // 192, 26,  0,  0
	NewHistoryState(26, 197, 26, 52, 25), // 193, 26,  0,  1
	NewHistoryState(26, 198, 52, 26, 26), // 194, 26,  1,  0
	NewHistoryState(26, 199, 53, 53, 26), // 195, 26,  1,  1
	NewHistoryState(27, 200, 25, 25, 25), // 196, 27,  0,  0
	NewHistoryState(27, 201, 26, 52, 25), // 197, 27,  0,  1
	NewHistoryState(27, 202, 52, 26, 26), // 198, 27,  1,  0
	NewHistoryState(27, 203, 53, 53, 26), // 199, 27,  1,  1
	NewHistoryState(28, 204, 25, 25, 25), // 200, 28,  0,  0
	NewHistoryState(28, 205, 26, 52, 25), // 201, 28,  0,  1
	NewHistoryState(28, 206, 52, 26, 26), // 202, 28,  1,  0
	NewHistoryState(28, 207, 53, 53, 26), // 203, 28,  1,  1
	NewHistoryState(29, 208, 25, 25, 25), // 204, 29,  0,  0
	NewHistoryState(29, 209, 26, 52, 25), // 205, 29,  0,  1
	NewHistoryState(29, 210, 52, 26, 26), // 206, 29,  1,  0
	NewHistoryState(29, 211, 53, 53, 26), // 207, 29,  1,  1
	NewHistoryState(30, 212, 25, 25, 25), // 208, 30,  0,  0
	NewHistoryState(30, 213, 26, 52, 25), // 209, 30,  0,  1
	NewHistoryState(30, 214, 52, 26, 26), // 210, 30,  1,  0
	NewHistoryState(30, 215, 53, 53, 26), // 211, 30,  1,  1
	NewHistoryState(31, 216, 25, 25, 25), // 212, 31,  0,  0
	NewHistoryState(31, 217, 26, 52, 25), // 213, 31,  0,  1
	NewHistoryState(31, 218, 52, 26, 26), // 214, 31,  1,  0
	NewHistoryState(31, 219, 53, 53, 26), // 215, 31,  1,  1
	NewHistoryState(32, 220, 25, 25, 25), // 216, 32,  0,  0
	NewHistoryState(32, 220, 26, 52, 25), // 217, 32,  0,  1
	NewHistoryState(32, 220, 52, 26, 26), // 218, 32,  1,  0
	NewHistoryState(32, 220, 53, 53, 26), // 219, 32,  1,  1
	NewHistoryState(33, 221, 53, 53, 26), // 220, 33,  1,  1
	NewHistoryState(34, 222, 53, 53, 26), // 221, 34,  1,  1
	NewHistoryState(35, 223, 53, 53, 26), // 222, 35,  1,  1
	NewHistoryState(36, 224, 53, 53, 26), // 223, 36,  1,  1
	NewHistoryState(37, 225, 53, 53, 26), // 224, 37,  1,  1
	NewHistoryState(38, 226, 53, 53, 26), // 225, 38,  1,  1
	NewHistoryState(39, 227, 53, 53, 26), // 226, 39,  1,  1
	NewHistoryState(40, 228, 53, 53, 26), // 227, 40,  1,  1
	NewHistoryState(41, 229, 53, 53, 26), // 228, 41,  1,  1
	NewHistoryState(42, 230, 53, 53, 26), // 229, 42,  1,  1
	NewHistoryState(43, 231, 53, 53, 26), // 230, 43,  1,  1
	NewHistoryState(44, 232, 53, 53, 26), // 231, 44,  1,  1
	NewHistoryState(45, 233, 53, 53, 26), // 232, 45,  1,  1
	NewHistoryState(46, 234, 53, 53, 26), // 233, 46,  1,  1
	NewHistoryState(47, 235, 53, 53, 26), // 234, 47,  1,  1
	NewHistoryState(48, 236, 53, 53, 26), // 235, 48,  1,  1
	NewHistoryState(49, 237, 53, 53, 26), // 236, 49,  1,  1
	NewHistoryState(50, 238, 53, 53, 26), // 237, 50,  1,  1
	NewHistoryState(51, 239, 53, 53, 26), // 238, 51,  1,  1
	NewHistoryState(52, 240, 53, 53, 26), // 239, 52,  1,  1
	NewHistoryState(53, 241, 53, 53, 26), // 240, 53,  1,  1
	NewHistoryState(54, 242, 53, 53, 26), // 241, 54,  1,  1
	NewHistoryState(55, 243, 53, 53, 26), // 242, 55,  1,  1
	NewHistoryState(56, 244, 53, 53, 26), // 243, 56,  1,  1
	NewHistoryState(57, 245, 53, 53, 26), // 244, 57,  1,  1
	NewHistoryState(58, 246, 53, 53, 26), // 245, 58,  1,  1
	NewHistoryState(59, 247, 53, 53, 26), // 246, 59,  1,  1
	NewHistoryState(60, 248, 53, 53, 26), // 247, 60,  1,  1
	NewHistoryState(61, 249, 53, 53, 26), // 248, 61,  1,  1
	NewHistoryState(62, 250, 53, 53, 26), // 249, 62,  1,  1
	NewHistoryState(63, 251, 53, 53, 26), // 250, 63,  1,  1
	NewHistoryState(64, 252, 53, 53, 26), // 251, 64,  1,  1
	NewHistoryState(65, 253, 53, 53, 26), // 252, 65,  1,  1
	NewHistoryState(66, 254, 53, 53, 26), // 253, 66,  1,  1
	NewHistoryState(67, 254, 53, 53, 26), // 254, 67,  1,  1
}
