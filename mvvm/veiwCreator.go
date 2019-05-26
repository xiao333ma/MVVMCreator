package mvvm

import (
	"./utils"
)

type ViewCreator struct{}

func (r ViewCreator) CreateDotHFile() {
	content := `//
//  __prefix__View.h
//  __proj__
//
//  Created by __user__ on __time__.
//  Copyright © __year__ __user__. All rights reserved.
//

#import "IGCBaseView.h"

NS_ASSUME_NONNULL_BEGIN

@interface __prefix__View : IGCBaseView

@end

NS_ASSUME_NONNULL_END
`

	utils.WriteFile(utils.FileTypeOfView, true, content)

}

func (r ViewCreator) CreateDotMFile() {
	content := `//
//  __prefix__View.m
//  __proj__
//
//  Created by __user__ on __time__.
//  Copyright © __year__ __user__. All rights reserved.
//

#import "__prefix__View.h"
#import "__prefix__Pipeline.h"

@interface __prefix__View ()

@property (nonatomic, strong) __prefix__Pipeline *pipeline;

@end

@implementation __prefix__View

@dynamic pipeline;

- (instancetype)initWithFrame:(CGRect)frame {
    self = [super initWithFrame:frame];
    if (self) {
        [self IGC_addSubViewAndLayout];
    }
    return self;
}

- (void)setupSubViewsPipline:(__kindof IGCPipeline *)rootPipeline {

}

- (void)addObservers {

}

#pragma mark - Layout

- (void)IGC_addSubViewAndLayout {

}

#pragma mark - Delegate

#pragma mark - Helper

#pragma mark - TargetAction

#pragma mark - Accessor

@end`

	utils.WriteFile(utils.FileTypeOfView, false, content)
}
