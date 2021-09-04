package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMatchStage(filter bson.D) bson.D {
	return bson.D{
		primitive.E{
			Key:   "$match",
			Value: filter,
		},
	}
}

func GetLimitStage(limit int) bson.D {
	return bson.D{
		primitive.E{
			Key:   "$limit",
			Value: limit,
		},
	}
}

func GetUnwindStage(path string, preserveNullAndEmptyArray bool) bson.D {
	return bson.D{
		primitive.E{
			Key: "$unwind",
			Value: bson.D{
				primitive.E{
					Key:   "path",
					Value: path,
				},
				primitive.E{
					Key:   "preserveNullAndEmptyArrays",
					Value: preserveNullAndEmptyArray,
				},
			},
		},
	}
}

func GetLookupStage(from string, localField string, foreignField string, as string) bson.D {
	return bson.D{
		primitive.E{
			Key: "$lookup",
			Value: bson.D{
				primitive.E{
					Key:   "from",
					Value: from,
				},
				primitive.E{
					Key:   "localField",
					Value: localField,
				},
				primitive.E{
					Key:   "foreignField",
					Value: foreignField,
				},
				primitive.E{
					Key:   "as",
					Value: as,
				},
			},
		},
	}
}
